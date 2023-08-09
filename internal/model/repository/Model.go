package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Model struct {
	Id        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
