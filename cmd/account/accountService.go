package account

type Account struct {
	UserName string `json:"userName"`
	Plan     Plan   `json: "Plan"`
}

var activeAccounts []Account

func getAll() *[]Account {

	activeAccounts = []Account{
		{UserName: "User1", Plan: Plan{Name: "Standard", Price: 15.5}},
		{UserName: "User2", Plan: Plan{Name: "Premium", Price: 25.5}},
	}

	return &activeAccounts
}

func addAccount(account *Account) {
	activeAccounts = append(activeAccounts, *account)
}
