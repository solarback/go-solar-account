package plan

import (
	"account-app/cmd/domain"
	. "account-app/internal/model/repository"
	"github.com/devfeel/mapper"
	"log"
)

type planRepo interface {
	GetById(id string) (Plan, error)
	GetActive() ([]Plan, error)
}

type Service struct {
	repository planRepo
}

func New(repository planRepo) *Service {
	return &Service{
		repository: repository,
	}
}

func (service *Service) GetById(id string) domain.Plan {
	plan, err := service.repository.GetById(id)

	if err != nil {
		log.Fatalf("Error on fetching plane by name=%s %v", id, err)
	}

	return toDomain(plan)
}

func (service *Service) GetActive() []domain.Plan {
	plan, err := service.repository.GetActive()

	if err != nil {
		log.Fatalf("Error on fetching active plans %v", err)
	}

	return toDomainSlice(plan)
}

func toDomain(from Plan) (target domain.Plan) {
	if err := mapper.Mapper(&from, &target); err != nil {
		log.Fatalf("Error on mapping plan %v", err)
	}
	return target
}

func toDomainSlice(from []Plan) (target []domain.Plan) {
	if err := mapper.Mapper(&from, &target); err != nil {
		log.Fatalf("Error on mapping plan %v", err)
	}
	return target
}
