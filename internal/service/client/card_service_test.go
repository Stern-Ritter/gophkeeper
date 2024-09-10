package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	pb "github.com/Stern-Ritter/gophkeeper/proto/gen/gophkeeper/gophkeeperapi/v1"
)

func TestCreateCard(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	cardClient := NewMockCardServiceV1Client(mockCtrl)
	cardClient.EXPECT().AddCard(gomock.Any(), &pb.AddCardRequestV1{
		Number:  "1234 5678 9012 3456",
		Owner:   "John Doe",
		Expiry:  "12/25",
		Cvc:     "123",
		Pin:     "1234",
		Comment: "comment",
	}).Return(&pb.AddCardResponseV1{}, nil)

	service := NewCardService(cardClient)
	err := service.CreateCard("1234 5678 9012 3456", "John Doe", "12/25", "123",
		"1234", "comment")

	assert.NoError(t, err, "Expected no error when creating card")
}

func TestGetAllCards(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	cardClient := NewMockCardServiceV1Client(mockCtrl)
	cardClient.EXPECT().GetCards(gomock.Any(), &pb.GetCardsRequestV1{}).
		Return(&pb.GetCardsResponseV1{
			Cards: []*pb.CardV1{
				{Number: "1234 5678 9012 3456", Owner: "John Doe", Expiry: "12/25", Comment: "comment 1"},
				{Number: "9876 5432 1098 7654", Owner: "Jane Doe", Expiry: "11/23", Comment: "comment 2"},
			},
		}, nil)

	service := NewCardService(cardClient)
	cards, err := service.GetAllCards()

	assert.NoError(t, err, "Expected no error when retrieving cards")
	assert.Len(t, cards, 2, "Expected two cards")

	assert.Equal(t, "1234 5678 9012 3456", cards[0].Number)
	assert.Equal(t, "John Doe", cards[0].Owner)
	assert.Equal(t, "12/25", cards[0].Expiry)
	assert.Equal(t, "comment 1", cards[0].Comment)

	assert.Equal(t, "9876 5432 1098 7654", cards[1].Number)
	assert.Equal(t, "Jane Doe", cards[1].Owner)
	assert.Equal(t, "11/23", cards[1].Expiry)
	assert.Equal(t, "comment 2", cards[1].Comment)
}

func TestDeleteCard(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	cardClient := NewMockCardServiceV1Client(mockCtrl)
	cardClient.EXPECT().DeleteCard(gomock.Any(), &pb.DeleteCardRequestV1{Id: "1"}).
		Return(&pb.DeleteCardResponseV1{}, nil)

	service := NewCardService(cardClient)
	err := service.DeleteCard("1")

	assert.NoError(t, err, "Expected no error when deleting card")
}
