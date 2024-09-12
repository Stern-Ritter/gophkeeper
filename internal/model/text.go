package model

import (
	pb "github.com/Stern-Ritter/gophkeeper/proto/gen/gophkeeper/gophkeeperapi/v1"
)

type Text struct {
	ID      string
	UserID  string
	Text    string
	Comment string
}

type SensitiveTextData struct {
	Text string `json:"text"`
}

func (t *Text) GetSensitiveTextData() SensitiveTextData {
	return SensitiveTextData{
		Text: t.Text,
	}
}

func (t *Text) SetSensitiveTextData(data SensitiveTextData) {
	t.Text = data.Text
}

func AddTextRequestToText(req *pb.AddTextRequestV1) Text {
	return Text{
		Text:    req.Text,
		Comment: req.Comment,
	}
}

func TextToTextMessage(t Text) *pb.TextV1 {
	return &pb.TextV1{
		Id:      t.ID,
		UserId:  t.UserID,
		Text:    t.Text,
		Comment: t.Comment,
	}
}

func TextsToRepeatedTextMessage(t []Text) []*pb.TextV1 {
	texts := make([]*pb.TextV1, len(t))
	for i, text := range t {
		texts[i] = TextToTextMessage(text)
	}

	return texts
}

func TextToData(t Text) Data {
	return Data{
		ID:      t.ID,
		UserID:  t.UserID,
		Type:    TextType,
		Comment: t.Comment,
	}
}

func DataToText(d Data) Text {
	return Text{
		ID:      d.ID,
		UserID:  d.UserID,
		Comment: d.Comment,
	}
}
