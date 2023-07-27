package dao

import (
	"account-app/cmd/dao/accountType"
	"account-app/cmd/dao/plan"
	"account-app/cmd/dao/user"
	"gorm.io/gorm"
)

type Repositories struct {
	UserRepo    *user.UserRepo
	PlanRepo    *plan.Repository
	AccountType *accountType.Repository
}

func InitRepositories(db *gorm.DB) *Repositories {
	var repos Repositories
	repos.UserRepo = user.New(db)
	repos.PlanRepo = plan.New(db)
	repos.AccountType = accountType.New(db)
	return &repos
}
