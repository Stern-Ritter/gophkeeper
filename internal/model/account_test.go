package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	pb "github.com/Stern-Ritter/gophkeeper/proto/gen/gophkeeper/gophkeeperapi/v1"
)

func TestGetSensitiveAccountData(t *testing.T) {
	account := &Account{
		Login:    "account login",
		Password: "account password",
	}

	sensitiveData := account.GetSensitiveAccountData()

	assert.Equal(t, "account login", sensitiveData.Login)
	assert.Equal(t, "account password", sensitiveData.Password)
}

func TestSetSensitiveAccountData(t *testing.T) {
	account := &Account{}
	data := SensitiveAccountData{
		Login:    "sensitive data login",
		Password: "sensitive data password",
	}

	account.SetSensitiveAccountData(data)

	assert.Equal(t, "sensitive data login", account.Login)
	assert.Equal(t, "sensitive data password", account.Password)
}

func TestAddAccountRequestToAccount(t *testing.T) {
	req := &pb.AddAccountRequestV1{
		Login:    "add account request login",
		Password: "add account request password",
		Comment:  "add account request comment",
	}

	account := AddAccountRequestToAccount(req)

	assert.Equal(t, "add account request login", account.Login)
	assert.Equal(t, "add account request password", account.Password)
	assert.Equal(t, "add account request comment", account.Comment)
}

func TestAccountToAccountMessage(t *testing.T) {
	account := Account{
		ID:       "account id",
		UserID:   "account user id",
		Login:    "account login",
		Password: "account password",
		Comment:  "account comment",
	}

	message := AccountToAccountMessage(account)

	assert.Equal(t, "account id", message.Id)
	assert.Equal(t, "account user id", message.UserId)
	assert.Equal(t, "account login", message.Login)
	assert.Equal(t, "account password", message.Password)
	assert.Equal(t, "account comment", message.Comment)
}

func TestAccountsToRepeatedAccountMessage(t *testing.T) {
	accounts := []Account{
		{
			ID:       "account id first",
			UserID:   "account user id first",
			Login:    "account login first",
			Password: "account password first",
			Comment:  "account comment first",
		},
		{
			ID:       "account id second",
			UserID:   "account user id second",
			Login:    "account login second",
			Password: "account password second",
			Comment:  "account comment second",
		},
	}

	messages := AccountsToRepeatedAccountMessage(accounts)

	require.Len(t, messages, 2)

	first := messages[0]
	assert.Equal(t, "account id first", first.Id)
	assert.Equal(t, "account user id first", first.UserId)
	assert.Equal(t, "account login first", first.Login)
	assert.Equal(t, "account password first", first.Password)
	assert.Equal(t, "account comment first", first.Comment)

	second := messages[1]
	assert.Equal(t, "account id second", second.Id)
	assert.Equal(t, "account user id second", second.UserId)
	assert.Equal(t, "account login second", second.Login)
	assert.Equal(t, "account password second", second.Password)
	assert.Equal(t, "account comment second", second.Comment)
}

func TestAccountToData(t *testing.T) {
	account := Account{
		ID:      "account id",
		UserID:  "account user id",
		Comment: "account comment",
	}

	data := AccountToData(account)

	assert.Equal(t, "account id", data.ID)
	assert.Equal(t, "account user id", data.UserID)
	assert.Equal(t, AccountType, data.Type)
	assert.Equal(t, "account comment", data.Comment)
}

func TestDataToAccount(t *testing.T) {
	data := Data{
		ID:      "data id",
		UserID:  "data user id",
		Comment: "data comment",
	}

	account := DataToAccount(data)

	assert.Equal(t, "data id", account.ID)
	assert.Equal(t, "data user id", account.UserID)
	assert.Equal(t, "data comment", account.Comment)
}
