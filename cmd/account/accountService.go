package account

import (
	"account-app/internal/model"
	"time"
)

var activeAccounts = []model.Account{
	{UserName: "User1", CreateDate: time.Now(), Plan: model.Plan{Name: "Standard", Price: 15.5}},
	{UserName: "User2", Plan: model.Plan{Name: "Premium", Price: 25.5}},
}

func getAll() *[]model.Account {
	return &activeAccounts
}

func addAccount(account *model.Account) {
	activeAccounts = append(activeAccounts, *account)
}
