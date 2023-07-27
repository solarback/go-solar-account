package domain

import (
	"account-app/pkg/utils"
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID
	UserName string
	Email    string
	Password string
	Accounts []Account
}

func Create(name string, pass string) User {
	return User{
		UserName: name,
		Password: utils.HashPassword(pass),
	}
}

func (user *User) AddEmail(email string) {
	//TODO add validation
	user.Email = email
}

func (user *User) AssignAccount(accountType string, planId string) {
	user.Accounts = append(user.Accounts, Account{TypeId: accountType, SubscriptionPlans: []SubscriptionPlan{{PlanId: planId}}})
}
