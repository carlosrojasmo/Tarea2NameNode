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
)

const (
	port = ":50055" //Quiza debamos usar distintos puertos segun en que trabajamos
	addressDataNode1  = "localhost:50051"
	addressDataNode2  = "localhost:50052"
	addressDataNode3  = "localhost:50053"
	addressDataNode4  = "localhost:50054"
)

var dataNodes = [4]string{addressDataNode1,addressDataNode2,addressDataNode3,addressDataNode4}

type server struct {
	pb.UnimplementedLibroServiceServer
}
var LibroChunks= make(map[string][]Chunk) //Nose si este diccionario funca bien

type Chunk struct{
	offset int
	direccion string
}

func (s* server) GetAddressChunks( book *pb.BookName,stream pb.LibroService_GetAddressChunksServer) error{
	bookNombre:=book.Name
	
	for _,chunkNecesario:=range LibroChunks[bookNombre]{
		partelibro:=pb.SendUbicacion{Ubicacion:chunkNecesario.direccion,Id:string(chunkNecesario.offset)}
		stream.Send(&partelibro)
	}
	return nil
}


func Log(ip string,nombreLibro string,parte int,cantidadPartes int ,primero bool) bool{
	log, err := os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY, 0600)
    if err != nil {
        panic(err)
    }
	defer log.Close()
	if primero==true{
		nuevaEntrada:=(nombreLibro+" "+string(cantidadPartes)+"\n"+string(parte)+" "+ip+"\n")
		if _, err = log.WriteString(nuevaEntrada); err != nil {
        	panic(err)
		}
	}else{
		nuevaEntrada:=(string(parte)+" "+ip+"\n")
		if _, err = log.WriteString(nuevaEntrada); err != nil {
        	panic(err)
		}
	}

	return true
}

func (s* server) SendPropuesta(ctx context.Context,prop *pb.Propuesta) (*pb.Propuesta, error) {
	fmt.Println("Analizando propuesta...")
	distribucion := []*pb.PropuestaChunk{}
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
		for _,badIp := range MaquinasCaidas{
			if badIp == dir {
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
	primera:=true
	Cantidad:=len(distribucion)
	for _,chunkIterativo := range distribucion{
		chunkOffset:=chunkIterativo.Offset
		chunkIP:=chunkIterativo.IpMaquina
		chunkBook:=chunkIterativo.NombreLibro
		Log(chunkIP,chunkBook,int(chunkOffset),Cantidad,primera)
		primera=false
		newChunk:=Chunk{offset:int(chunkOffset),direccion:chunkIP}
		LibroChunks[chunkBook]=append(LibroChunks[chunkBook],newChunk)
	}
	return &pb.Propuesta{Chunk : distribucion}, nil
}

func main() { 
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
