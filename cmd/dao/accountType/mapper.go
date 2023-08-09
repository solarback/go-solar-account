package accountType

import (
	"account-app/cmd/dao/plan"
	"account-app/cmd/domain"
	"account-app/internal/model/repository"
)

func MapToRepo(src domain.Account) (tgt repository.Account) {
	tgt.AccountTypeId = src.AccountTypeId
	if src.SubscriptionPlans != nil {
		for _, subPlan := range src.SubscriptionPlans {
			tgt.SubscriptionPlans = append(tgt.SubscriptionPlans, plan.MapToRepo(subPlan))
		}
	}
	return
}

func MapToDomain(src repository.Account) (tgt domain.Account) {
	tgt.AccountTypeId = src.AccountTypeId
	if src.SubscriptionPlans != nil {
		for _, subPlan := range src.SubscriptionPlans {
			tgt.SubscriptionPlans = append(tgt.SubscriptionPlans, plan.MapToDomain(subPlan))
		}
	}
	return
}

func MapToSliceRepo(src []domain.Account) (tgt []repository.Account) {
	for _, acc := range src {
		tgt = append(tgt, MapToRepo(acc))
	}
	return
}

func MapToSliceDomain(src []repository.Account) (tgt []domain.Account) {
	for _, acc := range src {
		tgt = append(tgt, MapToDomain(acc))
	}
	return
}
