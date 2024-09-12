package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	pb "github.com/Stern-Ritter/gophkeeper/proto/gen/gophkeeper/gophkeeperapi/v1"
)

func TestCreateText(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	textClient := NewMockTextServiceV1Client(mockCtrl)
	textClient.EXPECT().AddText(gomock.Any(), &pb.AddTextRequestV1{
		Text:    "note",
		Comment: "comment",
	}).Return(&pb.AddTextResponseV1{}, nil)

	service := NewTextService(textClient)
	err := service.CreateText("note", "comment")

	assert.NoError(t, err, "Expected no error when creating text")
}

func TestGetAllTexts(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	textClient := NewMockTextServiceV1Client(mockCtrl)
	textClient.EXPECT().GetTexts(gomock.Any(), &pb.GetTextsRequestV1{}).
		Return(&pb.GetTextsResponseV1{
			Texts: []*pb.TextV1{
				{Text: "note 1", Comment: "comment 1"},
				{Text: "note 2", Comment: "comment 2"},
			},
		}, nil)

	service := NewTextService(textClient)
	texts, err := service.GetAllTexts()

	assert.NoError(t, err, "Expected no error when retrieving texts")
	assert.Len(t, texts, 2, "Expected two texts")

	assert.Equal(t, "note 1", texts[0].Text)
	assert.Equal(t, "comment 1", texts[0].Comment)

	assert.Equal(t, "note 2", texts[1].Text)
	assert.Equal(t, "comment 2", texts[1].Comment)
}

func TestDeleteText(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	textClient := NewMockTextServiceV1Client(mockCtrl)
	textClient.EXPECT().DeleteText(gomock.Any(), &pb.DeleteTextRequestV1{Id: "1"}).
		Return(&pb.DeleteTextResponseV1{}, nil)

	service := NewTextService(textClient)
	err := service.DeleteText("1")

	assert.NoError(t, err, "Expected no error when deleting text")
}
