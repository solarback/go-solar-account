package model

import "time"

type Account struct {
	Id         string    `json:"id"`
	UserName   string    `json:"userName"`
	Plan       Plan      `json:"activePlan"`
	CreateDate time.Time `json:"createDate"`
}
