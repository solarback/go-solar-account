package repository

import (
	"time"
)

type SubscriptionPlan struct {
	Model
	Accounts  []Account `gorm:"many2many:account_plans;"`
	PlanId    string    `gorm:"type:varchar(32)"`
	Plan      Plan      `gorm:"foreignKey:PlanId"`
	StartDate time.Time
	EndDate   time.Time
}
