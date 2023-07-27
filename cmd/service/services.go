package service

import (
	"account-app/cmd/dao"
	"account-app/cmd/service/accountType"
	"account-app/cmd/service/plan"
	"account-app/cmd/service/user"
)

type Services struct {
	UserService        *user.Service
	PlanService        *plan.Service
	AccountTypeService *accountType.Service
}

func Init(repositories *dao.Repositories) *Services {
	var services Services
	services.UserService = user.New(repositories.UserRepo)
	services.PlanService = plan.New(repositories.PlanRepo)
	services.AccountTypeService = accountType.New(repositories.AccountType)
	return &services
}
