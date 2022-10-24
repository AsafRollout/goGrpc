package main

import (
	"log"
	"net"

	"goGrpc/speakers"

	"google.golang.org/grpc"
)

func main() {
	log.Println("starting")
	lis, err := net.Listen("tcp", ":9876")
	if err != nil {
		log.Printf("error %v\n", err)
	}
	log.Println("opening net")
	sss := speakers.Server{}
	grpcServer := grpc.NewServer()
	speakers.RegisterAudioOutputServer(grpcServer, &sss)
	log.Println("listening")
	err2 := grpcServer.Serve(lis)
	if err2 != nil {
		log.Printf("error %v\n", err2)
	}
}

