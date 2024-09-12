package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	pb "github.com/Stern-Ritter/gophkeeper/proto/gen/gophkeeper/gophkeeperapi/v1"
)

func TestGetSensitiveCardData(t *testing.T) {
	card := &Card{
		Number: "1111222233334444",
		Owner:  "owner",
		Expiry: "12/27",
		CVC:    "123",
		Pin:    "5678",
	}

	sensitiveData := card.GetSensitiveCardData()

	assert.Equal(t, "1111222233334444", sensitiveData.Number)
	assert.Equal(t, "owner", sensitiveData.Owner)
	assert.Equal(t, "12/27", sensitiveData.Expiry)
	assert.Equal(t, "123", sensitiveData.CVC)
	assert.Equal(t, "5678", sensitiveData.Pin)
}

func TestSetSensitiveCardData(t *testing.T) {
	card := &Card{}
	data := SensitiveCardData{
		Number: "1111222233334444",
		Owner:  "owner",
		Expiry: "11/25",
		CVC:    "321",
		Pin:    "8765",
	}

	card.SetSensitiveCardData(data)

	assert.Equal(t, "1111222233334444", card.Number)
	assert.Equal(t, "owner", card.Owner)
	assert.Equal(t, "11/25", card.Expiry)
	assert.Equal(t, "321", card.CVC)
	assert.Equal(t, "8765", card.Pin)
}

func TestAddCardRequestToCard(t *testing.T) {
	req := &pb.AddCardRequestV1{
		Number:  "1111222233334444",
		Owner:   "request owner",
		Expiry:  "01/23",
		Cvc:     "456",
		Pin:     "2345",
		Comment: "request comment",
	}

	card := AddCardRequestToCard(req)

	assert.Equal(t, "1111222233334444", card.Number)
	assert.Equal(t, "request owner", card.Owner)
	assert.Equal(t, "01/23", card.Expiry)
	assert.Equal(t, "456", card.CVC)
	assert.Equal(t, "2345", card.Pin)
	assert.Equal(t, "request comment", card.Comment)
}

func TestCardToCardMessage(t *testing.T) {
	card := Card{
		ID:      "1",
		UserID:  "42",
		Number:  "1111222233334444",
		Owner:   "owner",
		Expiry:  "11/25",
		CVC:     "789",
		Pin:     "3456",
		Comment: "card comment",
	}

	message := CardToCardMessage(card)

	assert.Equal(t, "1", message.Id)
	assert.Equal(t, "42", message.UserId)
	assert.Equal(t, "1111222233334444", message.Number)
	assert.Equal(t, "owner", message.Owner)
	assert.Equal(t, "11/25", message.Expiry)
	assert.Equal(t, "789", message.Cvc)
	assert.Equal(t, "3456", message.Pin)
	assert.Equal(t, "card comment", message.Comment)
}

func TestCardsToRepeatedCardMessage(t *testing.T) {
	cards := []Card{
		{
			ID:      "1",
			UserID:  "1",
			Number:  "1111222233334444",
			Owner:   "first",
			Expiry:  "12/24",
			CVC:     "123",
			Pin:     "5678",
			Comment: "comment first",
		},
		{
			ID:      "2",
			UserID:  "2",
			Number:  "1111222233334444",
			Owner:   "second",
			Expiry:  "01/25",
			CVC:     "456",
			Pin:     "8765",
			Comment: "comment second",
		},
	}

	messages := CardsToRepeatedCardMessage(cards)

	require.Len(t, messages, 2)

	first := messages[0]
	assert.Equal(t, "1", first.Id)
	assert.Equal(t, "1", first.UserId)
	assert.Equal(t, "1111222233334444", first.Number)
	assert.Equal(t, "first", first.Owner)
	assert.Equal(t, "12/24", first.Expiry)
	assert.Equal(t, "123", first.Cvc)
	assert.Equal(t, "5678", first.Pin)
	assert.Equal(t, "comment first", first.Comment)

	second := messages[1]
	assert.Equal(t, "2", second.Id)
	assert.Equal(t, "2", second.UserId)
	assert.Equal(t, "1111222233334444", second.Number)
	assert.Equal(t, "second", second.Owner)
	assert.Equal(t, "01/25", second.Expiry)
	assert.Equal(t, "456", second.Cvc)
	assert.Equal(t, "8765", second.Pin)
	assert.Equal(t, "comment second", second.Comment)
}

func TestCardToData(t *testing.T) {
	card := Card{
		ID:      "1",
		UserID:  "2",
		Comment: "card comment",
	}

	data := CardToData(card)

	assert.Equal(t, "1", data.ID)
	assert.Equal(t, "2", data.UserID)
	assert.Equal(t, CardType, data.Type)
	assert.Equal(t, "card comment", data.Comment)
}

func TestDataToCard(t *testing.T) {
	data := Data{
		ID:      "1",
		UserID:  "2",
		Comment: "data comment",
	}

	card := DataToCard(data)

	assert.Equal(t, "1", card.ID)
	assert.Equal(t, "2", card.UserID)
	assert.Equal(t, "data comment", card.Comment)
}
