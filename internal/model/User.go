package model

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `json:"id" mapper:"Id"`
	UserName string    `json:"userName" mapper:"Name"`
	Email    string    `json:"email" mapper:"Email"`
	Accounts []Account `json:"accounts" mapper:"Accounts"`
}
