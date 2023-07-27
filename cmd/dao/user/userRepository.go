package user

import (
	. "account-app/internal/model/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepo struct {
	conn *gorm.DB
}

func New(db *gorm.DB) *UserRepo {
	return &UserRepo{
		conn: db,
	}
}

func (repo *UserRepo) GetById(id uuid.UUID) (User, error) {
	var user User
	result := repo.conn.Where(&user, id)
	return user, result.Error
}

func (repo *UserRepo) GetAll() ([]User, error) {
	var users []User
	result := repo.conn.Find(&users)
	return users, result.Error
}

func (repo *UserRepo) AddUser(newUser User) (User, error) {
	result := repo.conn.Create(newUser)
	return newUser, result.Error
}
