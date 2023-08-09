package plan

import (
	"account-app/cmd/domain"
	"account-app/internal/model/repository"
)

func MapToRepo(src domain.SubscriptionPlan) (tgt repository.SubscriptionPlan) {
	tgt.PlanId = src.PlanId
	tgt.StartDate = src.StartDate
	return
}

func MapToDomain(src repository.SubscriptionPlan) (tgt domain.SubscriptionPlan) {
	tgt.PlanId = src.PlanId
	tgt.StartDate = src.StartDate
	return
}
