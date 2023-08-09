package model

type AccountType struct {
	Type  string `json:"value" mapper:"Type"`
	Title string `json:"label" mapper:"Title"`
}
