package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "leoho.io/grpc-go-example/proto"
)

func main() {
	connect, err := grpc.Dial(
		":50051",
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer connect.Close()

	client := pb.NewAlbumServiceClient(connect)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// createAlbum(ctx, client)
	// getAlbum("2a897be8-2c7c-49bd-b05a-d92a8a1cbfbc", ctx, client)
	getAlbums(ctx, client)
}

func createAlbum(ctx context.Context, client pb.AlbumServiceClient) {
	response, err := client.CreateAlbum(ctx, &pb.CreateAlbumRequest{
		Title:       "RINGO",
		Artist:      "ITZY",
		ReleaseDate: "2023/10/18",
	})
	if err != nil {
		log.Fatalf("could not create album: %v", err)
	}
	log.Printf("CreateAlbum response: %v", response)
}

func getAlbum(id string, ctx context.Context, client pb.AlbumServiceClient) {
	response, err := client.GetAlbum(ctx, &pb.GetAlbumRequest{
		Id: id,
	})
	if err != nil {
		log.Fatalf("could not get album: %v", err)
	}
	log.Printf("GetAlbum response: %v", response)
}

func getAlbums(ctx context.Context, client pb.AlbumServiceClient) {
	response, err := client.GetAlbums(ctx, &pb.GetAlbumsRequest{})
	if err != nil {
		log.Fatalf("could not get albums: %v", err)
	}
	log.Printf("GetAlbums response: %v", response)
}
