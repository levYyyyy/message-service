package main

import (
	"log"
	"net"

	k "github.com/levYyyyy/message-microservice/internal/kafka"
	messageService "github.com/levYyyyy/message-microservice/internal/message/proto"
	"github.com/levYyyyy/message-microservice/internal/server"
	"google.golang.org/grpc"
)

func main() {

	go k.StartKafka()

	// Creat gRPC server
	s := grpc.NewServer()
	srv := &server.MessageServer{}
	messageService.RegisterMessageServiceServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
