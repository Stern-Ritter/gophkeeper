package client

import (
	"testing"

	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	config "github.com/Stern-Ritter/gophkeeper/internal/config/client"
	pb "github.com/Stern-Ritter/gophkeeper/proto/gen/gophkeeper/gophkeeperapi/v1"
)

func TestDataView(t *testing.T) {
	app := tview.NewApplication()

	cfg := &config.ClientConfig{}

	client := NewClient(cfg)
	client.SetApp(app)

	element := client.DataView()

	menu, ok := element.(*tview.List)
	assert.True(t, ok, "Should create a menu")
	assert.Equal(t, 5, menu.GetItemCount(), "Menu should have 5 items")

	viewAccountsItem, _ := menu.GetItemText(0)
	assert.Equal(t, "View accounts data: login password pairs", viewAccountsItem, "First menu item should be 'View accounts data'")

	viewCardsItem, _ := menu.GetItemText(1)
	assert.Equal(t, "View cards data: payment cards details", viewCardsItem, "Second menu item should be 'View cards data'")

	viewTextItem, _ := menu.GetItemText(2)
	assert.Equal(t, "View text data: text notes", viewTextItem, "Third menu item should be 'View text data'")

	viewFilesItem, _ := menu.GetItemText(3)
	assert.Equal(t, "View file data: uploaded files", viewFilesItem, "Fourth menu item should be 'View file data'")

	backItem, _ := menu.GetItemText(4)
	assert.Equal(t, "Back", backItem, "Fifth menu item should be 'Back'")
}

func TestAccountsView(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	accountService := NewMockAccountService(mockCtrl)
	accountService.EXPECT().GetAllAccounts().Return([]*pb.AccountV1{
		{Id: "1", Login: "user 1", Password: "password 1", Comment: "comment 1"},
		{Id: "2", Login: "user 2", Password: "password 2", Comment: "comment 2"},
	}, nil)

	app := tview.NewApplication()
	cfg := &config.ClientConfig{}
	client := ClientImpl{
		accountService: accountService,
		config:         cfg,
	}
	client.SetApp(app)

	previousView := tview.NewTextView()
	element := client.AccountsView(previousView)

	table, ok := element.(*tview.Table)
	assert.True(t, ok, "Should create a table")
	assert.Equal(t, 3, table.GetRowCount(), "Table should have 3 rows (header and 2 data rows)")

	firstLoginCell := table.GetCell(1, 0).Text
	assert.Equal(t, "user 1", firstLoginCell, "First row, login column should contain 'user 1'")
	firstPasswordCell := table.GetCell(1, 1).Text
	assert.Equal(t, "password 1", firstPasswordCell, "First row, password column should contain 'password 1'")
	firstCommentCell := table.GetCell(1, 2).Text
	assert.Equal(t, "comment 1", firstCommentCell, "First row, comment column should contain 'comment 1'")

	secondLoginCell := table.GetCell(2, 0).Text
	assert.Equal(t, "user 2", secondLoginCell, "Second row, login column should contain 'user 2'")
	secondPasswordCell := table.GetCell(2, 1).Text
	assert.Equal(t, "password 2", secondPasswordCell, "Second row, password column should contain 'password 2'")
	secondCommentCell := table.GetCell(2, 2).Text
	assert.Equal(t, "comment 2", secondCommentCell, "Second row, comment column should contain 'comment 2'")

	firstDeleteCell := table.GetCell(1, 3).Text
	assert.Equal(t, deleteBtnText, firstDeleteCell, "Delete button text should match")
	secondDeleteCell := table.GetCell(2, 3).Text
	assert.Equal(t, deleteBtnText, secondDeleteCell, "Delete button text should match")
}

func TestCardsView(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	cardService := NewMockCardService(mockCtrl)
	cardService.EXPECT().GetAllCards().Return([]*pb.CardV1{
		{Id: "1", Number: "1234", Owner: "John Doe", Expiry: "12/25", Cvc: "123", Pin: "0000", Comment: "test 1"},
		{Id: "2", Number: "5678", Owner: "Jane Doe", Expiry: "11/24", Cvc: "456", Pin: "1111", Comment: "test 2"},
	}, nil)

	app := tview.NewApplication()
	cfg := &config.ClientConfig{}
	client := ClientImpl{
		cardService: cardService,
		config:      cfg,
	}
	client.SetApp(app)

	previousView := tview.NewTextView()
	element := client.CardsView(previousView)

	table, ok := element.(*tview.Table)
	assert.True(t, ok, "Should create a table")
	assert.Equal(t, 3, table.GetRowCount(), "Table should have 3 rows (header and 2 data rows)")

	firstNumberCell := table.GetCell(1, 0).Text
	assert.Equal(t, "1234", firstNumberCell, "First row, number column should contain '1234'")
	firstOwnerCell := table.GetCell(1, 1).Text
	assert.Equal(t, "John Doe", firstOwnerCell, "First row, owner column should contain 'John Doe'")
	firstExpiryCell := table.GetCell(1, 2).Text
	assert.Equal(t, "12/25", firstExpiryCell, "First row, expiry column should contain '12/25'")
	firstCvcCell := table.GetCell(1, 3).Text
	assert.Equal(t, "123", firstCvcCell, "First row, cvc column should contain '123'")
	firstPinCell := table.GetCell(1, 4).Text
	assert.Equal(t, "0000", firstPinCell, "First row, pin column should contain '0000'")
	firstCommentCell := table.GetCell(1, 5).Text
	assert.Equal(t, "test 1", firstCommentCell, "First row, comment column should contain 'test 1'")

	secondNumberCell := table.GetCell(2, 0).Text
	assert.Equal(t, "5678", secondNumberCell, "Second row, number column should contain '5678'")
	secondOwnerCell := table.GetCell(2, 1).Text
	assert.Equal(t, "Jane Doe", secondOwnerCell, "Second row, owner column should contain 'Jane Doe'")
	secondExpiryCell := table.GetCell(2, 2).Text
	assert.Equal(t, "11/24", secondExpiryCell, "Second row, expiry column should contain '11/24'")
	secondCvcCell := table.GetCell(2, 3).Text
	assert.Equal(t, "456", secondCvcCell, "Second row, cvc column should contain '456'")
	secondPinCell := table.GetCell(2, 4).Text
	assert.Equal(t, "1111", secondPinCell, "Second row, pin column should contain '1111'")
	secondCommentCell := table.GetCell(2, 5).Text
	assert.Equal(t, "test 2", secondCommentCell, "Second row, comment column should contain 'test 2'")

	firstDeleteCell := table.GetCell(1, 6).Text
	assert.Equal(t, deleteBtnText, firstDeleteCell, "Delete button text should match")
	secondDeleteCell := table.GetCell(2, 6).Text
	assert.Equal(t, deleteBtnText, secondDeleteCell, "Delete button text should match")
}

func TestTextsView(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	textService := NewMockTextService(mockCtrl)
	textService.EXPECT().GetAllTexts().Return([]*pb.TextV1{
		{Id: "1", Text: "note 1", Comment: "comment 1"},
		{Id: "2", Text: "note 2", Comment: "comment 2"},
	}, nil)

	app := tview.NewApplication()
	cfg := &config.ClientConfig{}
	client := ClientImpl{
		textService: textService,
		config:      cfg,
	}
	client.SetApp(app)

	previousView := tview.NewTextView()
	element := client.TextsView(previousView)

	table, ok := element.(*tview.Table)
	assert.True(t, ok, "Should create a table")
	assert.Equal(t, 3, table.GetRowCount(), "Table should have 3 rows (header and 2 data rows)")

	firstTextCell := table.GetCell(1, 0).Text
	assert.Equal(t, "note 1", firstTextCell, "First row, text column should contain 'note 1'")
	firstCommentCell := table.GetCell(1, 1).Text
	assert.Equal(t, "comment 1", firstCommentCell, "First row, comment column should contain 'comment 1'")

	secondTextCell := table.GetCell(2, 0).Text
	assert.Equal(t, "note 2", secondTextCell, "Second row, text column should contain 'note 2'")
	secondCommentCell := table.GetCell(2, 1).Text
	assert.Equal(t, "comment 2", secondCommentCell, "Second row, comment column should contain 'comment 2'")

	firstDeleteCell := table.GetCell(1, 2).Text
	assert.Equal(t, deleteBtnText, firstDeleteCell, "Delete button text should match")
	secondDeleteCell := table.GetCell(2, 2).Text
	assert.Equal(t, deleteBtnText, secondDeleteCell, "Delete button text should match")
}

func TestFilesView(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	fileService := NewMockFileService(mockCtrl)
	fileService.EXPECT().GetAllFiles().Return([]*pb.FileV1{
		{Id: "1", Name: "file1.txt", Size: 1024, Comment: "test file 1"},
		{Id: "2", Name: "file2.pdf", Size: 2048, Comment: "test file 2"},
	}, nil)

	app := tview.NewApplication()
	cfg := &config.ClientConfig{}
	client := ClientImpl{
		fileService: fileService,
		config:      cfg,
	}
	client.SetApp(app)

	previousView := tview.NewTextView()
	element := client.FilesView(previousView)

	table, ok := element.(*tview.Table)
	assert.True(t, ok, "Should create a table")
	assert.Equal(t, 3, table.GetRowCount(), "Table should have 3 rows (header and 2 data rows)")

	firstNameCell := table.GetCell(1, 0).Text
	assert.Equal(t, "file1.txt", firstNameCell, "First row, name column should contain 'file1.txt'")
	firstSizeCell := table.GetCell(1, 1).Text
	assert.Equal(t, "1.00 KB", firstSizeCell, "First row, size column should contain '1.00 KB'")
	firstCommentCell := table.GetCell(1, 2).Text
	assert.Equal(t, "test file 1", firstCommentCell, "First row, comment column should contain 'test file 1'")

	secondNameCell := table.GetCell(2, 0).Text
	assert.Equal(t, "file2.pdf", secondNameCell, "Second row, name column should contain 'file2.pdf'")
	secondSizeCell := table.GetCell(2, 1).Text
	assert.Equal(t, "2.00 KB", secondSizeCell, "Second row, size column should contain '2.00 KB'")
	secondCommentCell := table.GetCell(2, 2).Text
	assert.Equal(t, "test file 2", secondCommentCell, "Second row, comment column should contain 'test file 2'")

	firstDownloadCell := table.GetCell(1, 3).Text
	assert.Equal(t, downloadBtnText, firstDownloadCell, "First row, download column should contain the download button text")
	firstDeleteCell := table.GetCell(1, 4).Text
	assert.Equal(t, deleteBtnText, firstDeleteCell, "First row, delete column should contain the delete button text")

	secondDownloadCell := table.GetCell(2, 3).Text
	assert.Equal(t, downloadBtnText, secondDownloadCell, "Second row, download column should contain the download button text")
	secondDeleteCell := table.GetCell(2, 4).Text
	assert.Equal(t, deleteBtnText, secondDeleteCell, "Second row, delete column should contain the delete button text")
}
