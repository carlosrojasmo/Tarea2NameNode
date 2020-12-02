package main

import (
	"io/ioutil"
	"log"
	"net"
	"context"
	"google.golang.org/grpc"
	pb "../proto"
	"os"
	"time"
	"math/rand"
	"fmt"
	"bufio"
	"strings"
)

const (
	port = ":50051" //Quiza debamos usar distintos puertos segun en que trabajamos
	addressDataNode1  = "localhost:50051"
	addressDataNode2  = "localhost:50051"
	addressDataNode3  = "localhost:50051"
)
var lockCentralizado = false
var dataNodes = [3]string{addressDataNode1,addressDataNode2,addressDataNode3}

var funcionamiento = "EsperaInput"

type server struct {
	pb.UnimplementedLibroServiceServer
}
var libroChunks= make(map[string][]aChunk) //Nose si este diccionario funca bien

type aChunk struct{
	offset int
	direccion string
}

func (s* server) GetAddressChunks( book *pb.BookName,stream pb.LibroService_GetAddressChunksServer) error{
	bookNombre:=book.Name
	
	for _,chunkNecesario:=range libroChunks[bookNombre]{
		partelibro:=pb.SendUbicacion{Ubicacion:chunkNecesario.direccion,Id:strings.Split(bookNombre,".")[0]+"--"+fmt.Sprint(chunkNecesario.offset)}
		stream.Send(&partelibro)
	}
	return nil
}

func (s* server) VerStatus2(ctx context.Context,prop *pb.Propuesta) (*pb.Status, error) {
	distribucion := prop.GetChunk()
	primera:=true
	Cantidad:=len(distribucion)
	for _,chunkIterativo := range distribucion{
		chunkOffset:=chunkIterativo.Offset
		chunkIP:=chunkIterativo.IpMaquina
		chunkBook:=chunkIterativo.NombreLibro
		escribirLog(chunkIP,chunkBook,int(chunkOffset),Cantidad,primera)
		primera=false
		newChunk:=aChunk{offset:int(chunkOffset),direccion:chunkIP}
		libroChunks[chunkBook]=append(libroChunks[chunkBook],newChunk)
	}

	return &pb.Status{Status:"ok"}, nil
}


func escribirLog(ip string,nombreLibro string,parte int,cantidadPartes int ,primero bool) bool{
	log, err := os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY, 0600)
    if err != nil {
        log,err=os.Create("log.txt")
    }
	defer log.Close()
	if primero==true{
		nuevaEntrada:=(nombreLibro+" "+fmt.Sprint(cantidadPartes)+"\n"+fmt.Sprint(parte)+" "+ip+"\n")
		if _, err = log.WriteString(nuevaEntrada); err != nil {
        	panic(err)
		}
	}else{
		nuevaEntrada:=(fmt.Sprint(parte)+" "+ip+"\n")
		if _, err = log.WriteString(nuevaEntrada); err != nil {
        	panic(err)
		}
	}

	return true
}

func (s* server) SendPropuesta(ctx context.Context,prop *pb.Propuesta) (*pb.Propuesta, error) {
	distribucion := []*pb.PropuestaChunk{}

	if funcionamiento == "C"{
	fmt.Println("Analizando propuesta...")
	
	fallos := []*pb.PropuestaChunk{}
	MaquinasCaidas := []string{}
	for _, chub := range prop.GetChunk() { //Revisamos el status de cada maquina
		ip := chub.GetIpMaquina()
		conn, err := grpc.Dial(ip, grpc.WithInsecure(), grpc.WithBlock())
    	if err != nil {
    		log.Fatalf("did not connect: %v", err)
    	}
    	defer conn.Close()
    	c := pb.NewLibroServiceClient(conn)
    	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		fmt.Println("Viendo status")
		status, err := c.VerStatus(ctx,&pb.Status{Status : ""})
		if status.GetStatus() != "ok" || err != nil {
			fallos = append(fallos,chub)
			MaquinasCaidas = append(MaquinasCaidas,ip)
		} else {
			distribucion = append(distribucion,chub)
		}
		fmt.Println("Status: ",status.GetStatus())

	}

	Nodes := []string{}

	for _,dir := range dataNodes{
		in := false
		for _,badIP := range MaquinasCaidas{
			if badIP == dir {
				in = true
				break
			}
		}

		if !in {
			Nodes = append(Nodes,dir)
		}
	}
    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)

	for _,badchub := range fallos {
		distribucion = append(distribucion,&pb.PropuestaChunk{Offset : badchub.GetOffset(),
			IpMaquina : Nodes[r1.Intn(len(Nodes))], NombreLibro : badchub.GetNombreLibro() })
	} 
	fmt.Println("Nueva distribucion enviada!")
    } else {
    	distribucion = prop.GetChunk()
    }
	primera:=true
	Cantidad:=len(distribucion)
	for _,chunkIterativo := range distribucion{
		chunkOffset:=chunkIterativo.Offset
		chunkIP:=chunkIterativo.IpMaquina
		chunkBook:=chunkIterativo.NombreLibro
		for{
			if lockCentralizado==false{
				break
			}
			time.Sleep(time.Second)
		}
		lockCentralizado=true
		escribirLog(chunkIP,chunkBook,int(chunkOffset),Cantidad,primera)
		lockCentralizado=false
		primera=false
		newChunk:=aChunk{offset:int(chunkOffset),direccion:chunkIP}
		libroChunks[chunkBook]=append(libroChunks[chunkBook],newChunk)
	}
	
	return &pb.Propuesta{Chunk : distribucion}, nil
}

func main() { 
	reader := bufio.NewReader(os.Stdin)
    fmt.Println("NameNode")
    fmt.Println("---------------------")

    
    fmt.Print("Indique si desea el funcionamiento centralizado o distribuido (C o D) : ")//se pide si es Centralizado o Distribuido
    input1, _ := reader.ReadString('\n')
    input1 = strings.Replace(input1, "\n", "", -1)
    input1 = strings.Replace(input1, "\r", "", -1)
    funcionamiento = input1
    if funcionamiento != "C" && funcionamiento != "D"{
    	log.Fatalf("Ingreso mal el tipo de funcionamiento ,abortando")
    }

	initial:=[]byte("")
	err := ioutil.WriteFile("log.txt", initial, 0644)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLibroServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
