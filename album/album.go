package album

import (
	"context"

	pb "leoho.io/grpc-go-example/proto"
)

func CreateAlbum(ctx context.Context, req *pb.CreateAlbumRequest) (result string, err error) {
	return "Album created", nil
}

func GetAlbum(ctx context.Context, req *pb.GetAlbumRequest) (result *pb.CreateAlbumRequest, err error) {
	return &pb.CreateAlbumRequest{
		Title:       "RINGO",
		Artist:      "ITZY",
		ReleaseDate: "2023/10/18",
	}, nil
}
