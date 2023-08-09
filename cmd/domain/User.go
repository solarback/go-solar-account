package domain

import (
	"account-app/pkg/utils"
	"fmt"
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

	hashed, err := utils.HashPassword(pass)
	if err != nil {
		panic(fmt.Sprintf("Can't hash password %v", err))
	}

	return User{
		UserName: name,
		Password: hashed,
	}
}

func (user *User) AddEmail(email string) {
	//TODO add validation
	user.Email = email
}

func (user *User) AssignAccountType(accountType string) {
	user.Accounts = append(user.Accounts, Account{AccountTypeId: accountType})
}

func (user *User) AssignAccount(accountType string, planId string) {
	user.Accounts = append(user.Accounts, Account{AccountTypeId: accountType, SubscriptionPlans: []SubscriptionPlan{{PlanId: planId}}})
}

func (user *User) CheckPassword(password string) bool {
	res, err := utils.PasswordMatches(password, user.Password)
	if err != nil {
		panic(fmt.Sprintf("Error on password check"))
	}

	return res
}
