package service

import (
	"account-app/cmd/dao"
	"account-app/cmd/service/accountType"
	"account-app/cmd/service/auth"
	"account-app/cmd/service/plan"
	"account-app/cmd/service/user"
)

type Services struct {
	UserService        *user.Service
	PlanService        *plan.Service
	AccountTypeService *accountType.Service
	SecurityService    *auth.Service
}

func Init(repositories *dao.Repositories, tokenService *auth.Auth) *Services {
	var services Services
	services.UserService = user.New(repositories.UserRepo)
	services.PlanService = plan.New(repositories.PlanRepo)
	services.AccountTypeService = accountType.New(repositories.AccountType)
	services.SecurityService = auth.New(repositories.UserRepo, tokenService)
	return &services
}
