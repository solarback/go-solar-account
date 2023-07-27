package repository

import (
	"github.com/google/uuid"
	"time"
)

type SubscriptionPlan struct {
	Model
	AccountRefer uuid.UUID `gorm:"index:,unique"`
	PlanId       string    `gorm:"type:varchar(32)"`
	Plan         Plan      `gorm:"foreignKey:PlanId"`
	StartDate    time.Time
	EndDate      time.Time
}
