package user

import (
	"account-app/cmd/domain"
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

func (repo *UserRepo) GetById(id uuid.UUID) (domain.User, error) {
	var user User
	result := repo.conn.Where(&user, id)
	return MapToDomain(user), result.Error
}

func (repo *UserRepo) GetByEmail(email string) (domain.User, error) {
	var user User
	result := repo.conn.Where(&user, "Email = ?", email)
	return MapToDomain(user), result.Error
}

func (repo *UserRepo) GetAll() ([]domain.User, error) {
	var users []User
	result := repo.conn.Find(&users)
	return MapSliceToDomain(users), result.Error
}
