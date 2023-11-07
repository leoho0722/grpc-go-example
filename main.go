package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	db "leoho.io/grpc-go-example/database"
	pb "leoho.io/grpc-go-example/proto"
	server "leoho.io/grpc-go-example/server"
)

func main() {
	db.ConnectDB()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAlbumServiceServer(s, &server.Server{})
	log.Printf("Server listening at %v", lis.Addr())
	reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
