package model

import (
	"github.com/google/uuid"
	"time"
)

type SubscriptionPlan struct {
	Id        uuid.UUID `json:"id"         gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	AccountId uuid.UUID `json:"accountId"  gorm:"type:uuid"`
	PlanId    uuid.UUID `json:"planId"     gorm:"type:uuid"`
	Plan      Plan      `json:"plan"       gorm:"foreignKey:PlanId"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}
