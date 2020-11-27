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
)

var dataNodes = [3]string{addressDataNode1,addressDataNode2,addressDataNode3}

type server struct {
	pb.UnimplementedLibroServiceServer
}
var LibroAux= []Chunk{}
var ListaLibros=[]libro{}
var LibroChunks= make(map[string][]Chunk) //Nose si este diccionario funca bien
type libro struct{
	nombre string
	cantidadChunks int
}
type Chunk struct{
	offset int
	direccion string
}

func Log(libro libro, chunk Chunk) bool{
	log, err := os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY, 0600)
    if err != nil {
        panic(err)
    }
	defer log.Close()
	nombreLibro:=libro.nombre
	cantidadPartes:=libro.cantidadChunks
	parte:=chunk.offset
	ip:=chunk.direccion
	nuevaEntrada:=(nombreLibro+" "+string(cantidadPartes)+"\n"+string(parte)+" "+ip+"\n")
	if _, err = log.WriteString(nuevaEntrada); err != nil {
        panic(err)
    }

	return true
}

func Ubicar() Chunk{
	//Debe hacer la ubicacion del libro, acordar quiza(?)
	NuevaDireccion:=Chunk{}
	return NuevaDireccion
}
//Primero cuando lleguen los chunks se acumularan en LibroAUX , luego al llegar al offset final pasara la siguiente funcion
func newLibro(nombre string,cantidadChunks int) libro{
	libroNuevo:=libro{nombre:nombre,cantidadChunks:cantidadChunks}
	ListaLibros=append(ListaLibros,libroNuevo)
	LibroChunks[libroNuevo.nombre]=LibroAux
	return libroNuevo
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