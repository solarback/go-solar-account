package model

type Account struct {
	UserName string `json:"userName"`
	Plan     Plan   `json: "Plan"`
}
