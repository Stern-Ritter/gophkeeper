package client

import (
	"context"

	pb "github.com/Stern-Ritter/gophkeeper/proto/gen/gophkeeper/gophkeeperapi/v1"
)

// TextService defines an interface for user texts data.
type TextService interface {
	CreateText(text string, comment string) error
	GetAllTexts() ([]*pb.TextV1, error)
	DeleteText(textID string) error
}

// TextServiceImpl is implementation of the TextService interface.
type TextServiceImpl struct {
	textClient pb.TextServiceV1Client
}

// NewTextService creates a new instance of TextService.
func NewTextService(textClient pb.TextServiceV1Client) TextService {
	return &TextServiceImpl{
		textClient: textClient,
	}
}

// CreateText creates new text data with the given text and comment by sending a request
// to the account service with gRPC.
// Returns an error if the text data creation fails.
func (s *TextServiceImpl) CreateText(text string, comment string) error {
	ctx := context.Background()
	req := &pb.AddTextRequestV1{
		Text:    text,
		Comment: comment,
	}

	_, err := s.textClient.AddText(ctx, req)
	return err
}

// GetAllTexts get all texts data by sending a request to the text service with gRPC.
func (s *TextServiceImpl) GetAllTexts() ([]*pb.TextV1, error) {
	ctx := context.Background()
	req := &pb.GetTextsRequestV1{}
	resp, err := s.textClient.GetTexts(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.Texts, nil
}

// DeleteText deletes text data identified by the given id by sending a request
// to the text service with gRPC.
// Returns an error if the text data deletion fails.
func (s *TextServiceImpl) DeleteText(textID string) error {
	ctx := context.Background()
	req := &pb.DeleteTextRequestV1{Id: textID}

	_, err := s.textClient.DeleteText(ctx, req)
	return err
}
