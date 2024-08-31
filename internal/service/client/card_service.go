package client

import (
	"context"

	pb "github.com/Stern-Ritter/gophkeeper/proto/gen/gophkeeper/gophkeeperapi/v1"
)

// CardService defines an interface for user cards data.
type CardService interface {
	CreateCard(number string, owner string, expiry string, cvc string, pin string, comment string) error
	GetAllCards() ([]*pb.CardV1, error)
	DeleteCard(cardID string) error
}

// CardServiceImpl is implementation of the CardService interface.
type CardServiceImpl struct {
	cardClient pb.CardServiceV1Client
}

// NewCardService creates a new instance of CardService.
func NewCardService(cardClient pb.CardServiceV1Client) CardService {
	return &CardServiceImpl{
		cardClient: cardClient,
	}
}

// CreateCard creates new card data with the given number, owner, expiry, cvc, pin and comment by sending a request
// to the card service with gRPC.
// Returns an error if the card data creation fails.
func (s *CardServiceImpl) CreateCard(number string, owner string, expiry string, cvc string, pin string, comment string) error {
	ctx := context.Background()
	req := &pb.AddCardRequestV1{
		Number:  number,
		Owner:   owner,
		Expiry:  expiry,
		Cvc:     cvc,
		Pin:     pin,
		Comment: comment,
	}

	_, err := s.cardClient.AddCard(ctx, req)
	return err
}

// GetAllCards get all cards data by sending a request to the card service with gRPC.
func (s *CardServiceImpl) GetAllCards() ([]*pb.CardV1, error) {
	ctx := context.Background()
	req := &pb.GetCardsRequestV1{}
	resp, err := s.cardClient.GetCards(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.Cards, nil
}

// DeleteCard deletes card data identified by the given id by sending a request
// to the card service with gRPC.
// Returns an error if the card data deletion fails.
func (s *CardServiceImpl) DeleteCard(cardID string) error {
	ctx := context.Background()
	req := &pb.DeleteCardRequestV1{Id: cardID}

	_, err := s.cardClient.DeleteCard(ctx, req)
	return err
}
