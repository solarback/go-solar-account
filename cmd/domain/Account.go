package domain

import (
	"github.com/google/uuid"
	"time"
)

type Account struct {
	Id                uuid.UUID
	TypeId            string
	SubscriptionPlans []SubscriptionPlan
	CreateDate        time.Time
}
