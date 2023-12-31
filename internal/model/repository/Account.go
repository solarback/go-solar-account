package repository

import "github.com/google/uuid"

type Account struct {
	Model
	UserId            uuid.UUID
	AccountTypeId     string
	Type              AccountType        `gorm:"foreignKey:AccountTypeId;not null"`
	SubscriptionPlans []SubscriptionPlan `gorm:"many2many:account_plans;"`
}
