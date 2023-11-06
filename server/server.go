package server

import (
	"context"

	album "leoho.io/grpc-go-example/album"
	pb "leoho.io/grpc-go-example/proto"
)

type Server struct {
	pb.UnimplementedAlbumServiceServer
}

func (s *Server) CreateAlbum(ctx context.Context, req *pb.CreateAlbumRequest) (*pb.AlbumGeneralResponse, error) {
	res, err := album.CreateAlbum(ctx, req)
	response := &pb.AlbumGeneralResponse{
		Message: res,
	}
	return response, err
}

func (s *Server) GetAlbum(ctx context.Context, req *pb.GetAlbumRequest) (*pb.GetAlbumResponse, error) {
	res, err := album.GetAlbum(ctx, req)
	response := &pb.GetAlbumResponse{
		Message: "Success",
		Album:   res,
	}
	return response, err
}
