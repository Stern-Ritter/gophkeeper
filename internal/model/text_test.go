package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	pb "github.com/Stern-Ritter/gophkeeper/proto/gen/gophkeeper/gophkeeperapi/v1"
)

func TestGetSensitiveTextData(t *testing.T) {
	text := &Text{
		Text: "sensitive text",
	}

	sensitiveData := text.GetSensitiveTextData()

	assert.Equal(t, "sensitive text", sensitiveData.Text)
}

func TestSetSensitiveTextData(t *testing.T) {
	text := &Text{}
	data := SensitiveTextData{
		Text: "sensitive text",
	}

	text.SetSensitiveTextData(data)

	assert.Equal(t, "sensitive text", text.Text)
}

func TestAddTextRequestToText(t *testing.T) {
	req := &pb.AddTextRequestV1{
		Text:    "add text request text",
		Comment: "add text request comment",
	}

	text := AddTextRequestToText(req)

	assert.Equal(t, "add text request text", text.Text)
	assert.Equal(t, "add text request comment", text.Comment)
}

func TestTextToTextMessage(t *testing.T) {
	text := Text{
		ID:      "1",
		UserID:  "2",
		Text:    "text",
		Comment: "comment",
	}

	message := TextToTextMessage(text)

	assert.Equal(t, "1", message.Id)
	assert.Equal(t, "2", message.UserId)
	assert.Equal(t, "text", message.Text)
	assert.Equal(t, "comment", message.Comment)
}

func TestTextsToRepeatedTextMessage(t *testing.T) {
	texts := []Text{
		{
			ID:      "1",
			UserID:  "1",
			Text:    "first text",
			Comment: "first comment",
		},
		{
			ID:      "2",
			UserID:  "2",
			Text:    "second text",
			Comment: "second comment",
		},
	}

	messages := TextsToRepeatedTextMessage(texts)

	require.Len(t, messages, 2)

	first := messages[0]
	assert.Equal(t, "1", first.Id)
	assert.Equal(t, "1", first.UserId)
	assert.Equal(t, "first text", first.Text)
	assert.Equal(t, "first comment", first.Comment)

	second := messages[1]
	assert.Equal(t, "2", second.Id)
	assert.Equal(t, "2", second.UserId)
	assert.Equal(t, "second text", second.Text)
	assert.Equal(t, "second comment", second.Comment)
}

func TestTextToData(t *testing.T) {
	text := Text{
		ID:      "1",
		UserID:  "2",
		Comment: "comment",
	}

	data := TextToData(text)

	assert.Equal(t, "1", data.ID)
	assert.Equal(t, "2", data.UserID)
	assert.Equal(t, TextType, data.Type)
	assert.Equal(t, "comment", data.Comment)
}

func TestDataToText(t *testing.T) {
	data := Data{
		ID:      "1",
		UserID:  "2",
		Comment: "comment",
	}

	text := DataToText(data)

	assert.Equal(t, "1", text.ID)
	assert.Equal(t, "2", text.UserID)
	assert.Equal(t, "comment", text.Comment)
}
