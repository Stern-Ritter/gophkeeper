package client

import (
	"testing"

	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestSelectView(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockApp := NewMockApplication(ctrl)

	view := tview.NewBox()
	mockApp.EXPECT().SetRoot(view, true).Return(nil).Times(1)
	mockApp.EXPECT().SetFocus(view).Times(1)

	client := &ClientImpl{app: mockApp}

	client.SelectView(view)
}

func TestUpdateDraw(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockApp := NewMockApplication(ctrl)
	mockApp.EXPECT().Draw().Times(1)

	client := &ClientImpl{app: mockApp}

	client.UpdateDraw()
}

func TestQueueUpdateDraw(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockApp := NewMockApplication(ctrl)
	updateFunc := func() {}

	mockApp.EXPECT().QueueUpdateDraw(gomock.Any()).Do(func(fn func()) { fn() }).Times(1)

	client := &ClientImpl{app: mockApp}

	client.QueueUpdateDraw(updateFunc)
}

func TestNewClickableCell(t *testing.T) {
	handlerCalled := false
	handler := func() { handlerCalled = true }

	cell := newClickableCell("Cell", handler)
	assert.Equal(t, "Cell", cell.Text, "unexpected clickable cell text")

	cell.Clicked()

	assert.True(t, handlerCalled, "handler should be called after click")
}

func TestStopApp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockApp := NewMockApplication(ctrl)

	mockApp.EXPECT().Stop().Times(1)

	stopApp(mockApp)
}
