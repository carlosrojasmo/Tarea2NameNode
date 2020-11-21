package main

import (
	"io/ioutil"
	"log"
	"net"
	"context"
	"google.golang.org/grpc"
	pb "../proto"
	"fmt"
	"os"
)

const (
	port = ":50051" //Quiza debamos usar distintos puertos segun en que trabajamos
)
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
