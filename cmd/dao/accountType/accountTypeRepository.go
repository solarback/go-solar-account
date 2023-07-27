package accountType

import (
	. "account-app/internal/model/repository"
	"gorm.io/gorm"
)

type Repository struct {
	conn *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		conn: db,
	}
}

func (repo *Repository) GetUserTypes() (availableTypes []AccountType, err error) {
	result := repo.conn.Where("is_client = ?", true).Find(&availableTypes)
	return availableTypes, result.Error
}

func (repo *Repository) GetAdminTypes() (availableTypes []AccountType, err error) {
	result := repo.conn.Where("is_client = ?", false).Find(&availableTypes)
	return availableTypes, result.Error
}
