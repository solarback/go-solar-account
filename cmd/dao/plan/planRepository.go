package plan

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

func (repo *Repository) GetById(id string) (Plan, error) {
	var plan Plan
	result := repo.conn.Where("Id = ?", id).First(&plan)
	return plan, result.Error
}

func (repo *Repository) GetActive() ([]Plan, error) {
	var plan []Plan
	result := repo.conn.Where("Active = ?", true).Find(&plan)
	return plan, result.Error
}
