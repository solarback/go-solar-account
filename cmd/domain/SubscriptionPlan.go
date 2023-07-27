package domain

import (
	"github.com/google/uuid"
	"time"
)

type SubscriptionPlan struct {
	Id        uuid.UUID
	AccountId uuid.UUID
	PlanId    string
	StartDate time.Time
	EndDate   time.Time
}
