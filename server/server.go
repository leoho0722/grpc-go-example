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
	return res, err
}

func (s *Server) GetAlbum(ctx context.Context, req *pb.GetAlbumRequest) (*pb.GetAlbumResponse, error) {
	res, err := album.GetAlbum(ctx, req)
	return res, err
}

func (s *Server) GetAlbums(ctx context.Context, req *pb.GetAlbumsRequest) (*pb.GetAlbumsResponse, error) {
	res, err := album.GetAlbums(ctx, req)
	return res, err
}
