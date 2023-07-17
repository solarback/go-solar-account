package account

import "account-app/internal/model"

var activeAccounts []model.Account

func getAll() *[]model.Account {

	activeAccounts = []model.Account{
		{UserName: "User1", Plan: model.Plan{Name: "Standard", Price: 15.5}},
		{UserName: "User2", Plan: model.Plan{Name: "Premium", Price: 25.5}},
	}

	return &activeAccounts
}

func addAccount(account *model.Account) {
	activeAccounts = append(activeAccounts, *account)
}
