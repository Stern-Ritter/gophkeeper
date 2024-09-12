package model

type DataType string

const (
	AccountType = DataType("ACCOUNT")
	CardType    = DataType("CARD")
	TextType    = DataType("TEXT")
)

type Data struct {
	ID            string
	UserID        string
	Type          DataType
	SensitiveData []byte
	Comment       string
}
