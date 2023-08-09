package user

import (
	"account-app/cmd/dao/accountType"
	"account-app/cmd/domain"
	"account-app/internal/model/repository"
	"github.com/devfeel/mapper"
)

func init() {
	mapper.Register(&repository.User{})
	mapper.Register(&domain.User{})
}

func MapToRepo(user domain.User) (repoUser repository.User) {
	repoUser.Name = user.UserName
	repoUser.Password = user.Password
	repoUser.Email = user.Email

	if user.Accounts != nil {
		repoUser.Accounts = accountType.MapToSliceRepo(user.Accounts)
	}
	return
}

func MapToDomain(repoUser repository.User) (user domain.User) {
	user.UserName = repoUser.Name
	user.Password = repoUser.Password
	user.Email = repoUser.Email

	if repoUser.Accounts != nil {
		user.Accounts = accountType.MapToSliceDomain(repoUser.Accounts)
	}
	return
}

func MapSliceToDomain(repoUsers []repository.User) (users []domain.User) {
	for _, user := range repoUsers {
		users = append(users, MapToDomain(user))
	}

	return
}
