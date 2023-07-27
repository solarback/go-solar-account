package model

import (
	"github.com/google/uuid"
	"time"
)

type Account struct {
	Id                uuid.UUID          `json:"id"`
	Type              string             `json:"type"`
	SubscriptionPlans []SubscriptionPlan `json:"subscriptionPlans"`
	CreateDate        time.Time          `json:"createDate"`
}
