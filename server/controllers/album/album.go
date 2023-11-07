package controllers

import (
	"context"

	uuid "github.com/google/uuid"
	db "leoho.io/grpc-go-example/database"
	pb "leoho.io/grpc-go-example/proto"
)

func CreateAlbum(ctx context.Context, req *pb.CreateAlbumRequest) (result *pb.AlbumGeneralResponse, err error) {
	err = db.Create(&db.Album{
		Id:          uuid.New().String(),
		Title:       req.Album.Title,
		Artist:      req.Album.Artist,
		ReleaseDate: req.Album.ReleaseDate,
	})
	if err != nil {
		return &pb.AlbumGeneralResponse{
			Message: "Album create failed",
		}, err
	} else {
		return &pb.AlbumGeneralResponse{
			Message: "Album create success",
		}, nil
	}
}

func GetAlbum(ctx context.Context, req *pb.GetAlbumRequest) (result *pb.GetAlbumResponse, err error) {
	album, err := db.Get(req.Id)
	if err != nil {
		return &pb.GetAlbumResponse{
			Album:   nil,
			Message: "Get album failed",
		}, err
	} else {
		return &pb.GetAlbumResponse{
			Message: "Get album success",
			Album: &pb.Album{
				Id:          &album.Id,
				Title:       album.Title,
				Artist:      album.Artist,
				ReleaseDate: album.ReleaseDate,
			},
		}, nil
	}
}

func GetAlbums(ctx context.Context, req *pb.GetAlbumsRequest) (result *pb.GetAlbumsResponse, err error) {
	albums, err := db.GetAll()
	if err != nil {
		return &pb.GetAlbumsResponse{
			Albums:  nil,
			Message: "Get albums failed",
		}, err
	} else {
		var albumResponses []*pb.Album
		for _, album := range albums {
			albumResponses = append(albumResponses, &pb.Album{
				Id:          &album.Id,
				Title:       album.Title,
				Artist:      album.Artist,
				ReleaseDate: album.ReleaseDate,
			})
		}
		return &pb.GetAlbumsResponse{
			Message: "Get albums success",
			Albums:  albumResponses,
		}, nil
	}
}
