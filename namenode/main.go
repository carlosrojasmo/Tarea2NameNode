package main

import (
	"log"
	"net"
	"context"
	"google.golang.org/grpc"
	pb "../proto"
	"fmt"
)

const (
	port = ":50051" //Quiza debamos usar distintos puertos segun en que trabajamos
)
type server struct {
	pb.UnimplementedOrdenServiceServer
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

func main() { 
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterOrdenServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
