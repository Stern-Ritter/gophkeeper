package model

import (
	pb "github.com/Stern-Ritter/gophkeeper/proto/gen/gophkeeper/gophkeeperapi/v1"
)

type Account struct {
	ID       string
	UserID   string
	Login    string
	Password string
	Comment  string
}

type SensitiveAccountData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (a *Account) GetSensitiveAccountData() SensitiveAccountData {
	return SensitiveAccountData{
		Login:    a.Login,
		Password: a.Password,
	}
}

func (a *Account) SetSensitiveAccountData(data SensitiveAccountData) {
	a.Login = data.Login
	a.Password = data.Password
}

func AddAccountRequestToAccount(req *pb.AddAccountRequestV1) Account {
	return Account{
		Login:    req.Login,
		Password: req.Password,
		Comment:  req.Comment,
	}
}

func AccountToAccountMessage(a Account) *pb.AccountV1 {
	return &pb.AccountV1{
		Id:       a.ID,
		UserId:   a.UserID,
		Login:    a.Login,
		Password: a.Password,
		Comment:  a.Comment,
	}
}

func AccountsToRepeatedAccountMessage(a []Account) []*pb.AccountV1 {
	accounts := make([]*pb.AccountV1, len(a))
	for i, account := range a {
		accounts[i] = AccountToAccountMessage(account)
	}

	return accounts
}

func AccountToData(a Account) Data {
	return Data{
		ID:      a.ID,
		UserID:  a.UserID,
		Type:    AccountType,
		Comment: a.Comment,
	}
}

func DataToAccount(d Data) Account {
	return Account{
		ID:      d.ID,
		UserID:  d.UserID,
		Comment: d.Comment,
	}
}
