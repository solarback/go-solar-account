package user

import (
	"account-app/cmd/domain"
)

func (repo *UserRepo) AddUser(newUser domain.User) (domain.User, error) {
	toSave := MapToRepo(newUser)
	result := repo.conn.Create(&toSave)
	return MapToDomain(toSave), result.Error
}
