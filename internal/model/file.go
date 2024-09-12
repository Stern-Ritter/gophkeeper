package model

import (
	pb "github.com/Stern-Ritter/gophkeeper/proto/gen/gophkeeper/gophkeeperapi/v1"
)

type File struct {
	ID      string
	UserID  string
	Name    string
	Size    int64
	Path    string
	Comment string
}

func FileToFileMessage(f File) *pb.FileV1 {
	return &pb.FileV1{
		Id:      f.ID,
		UserId:  f.UserID,
		Name:    f.Name,
		Size:    f.Size,
		Comment: f.Comment,
	}
}

func FilesToRepeatedFileMessage(f []File) []*pb.FileV1 {
	files := make([]*pb.FileV1, len(f))
	for i, file := range f {
		files[i] = FileToFileMessage(file)
	}

	return files
}
