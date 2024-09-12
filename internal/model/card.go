package model

import (
	pb "github.com/Stern-Ritter/gophkeeper/proto/gen/gophkeeper/gophkeeperapi/v1"
)

type Card struct {
	ID      string
	UserID  string
	Number  string
	Owner   string
	Expiry  string
	CVC     string
	Pin     string
	Comment string
}

type SensitiveCardData struct {
	Number string `json:"number"`
	Owner  string `json:"owner"`
	Expiry string `json:"expiry"`
	CVC    string `json:"cvc"`
	Pin    string `json:"pin"`
}

func (c *Card) GetSensitiveCardData() SensitiveCardData {
	return SensitiveCardData{
		Number: c.Number,
		Owner:  c.Owner,
		Expiry: c.Expiry,
		CVC:    c.CVC,
		Pin:    c.Pin,
	}
}

func (c *Card) SetSensitiveCardData(data SensitiveCardData) {
	c.Number = data.Number
	c.Owner = data.Owner
	c.Expiry = data.Expiry
	c.CVC = data.CVC
	c.Pin = data.Pin
}

func AddCardRequestToCard(req *pb.AddCardRequestV1) Card {
	return Card{
		Number:  req.Number,
		Owner:   req.Owner,
		Expiry:  req.Expiry,
		CVC:     req.Cvc,
		Pin:     req.Pin,
		Comment: req.Comment,
	}
}

func CardToCardMessage(c Card) *pb.CardV1 {
	return &pb.CardV1{
		Id:      c.ID,
		UserId:  c.UserID,
		Number:  c.Number,
		Owner:   c.Owner,
		Expiry:  c.Expiry,
		Cvc:     c.CVC,
		Pin:     c.Pin,
		Comment: c.Comment,
	}
}

func CardsToRepeatedCardMessage(c []Card) []*pb.CardV1 {
	cards := make([]*pb.CardV1, len(c))
	for i, card := range c {
		cards[i] = CardToCardMessage(card)
	}

	return cards
}

func CardToData(c Card) Data {
	return Data{
		ID:      c.ID,
		UserID:  c.UserID,
		Type:    CardType,
		Comment: c.Comment,
	}
}

func DataToCard(d Data) Card {
	return Card{
		ID:      d.ID,
		UserID:  d.UserID,
		Comment: d.Comment,
	}
}
