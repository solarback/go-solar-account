package accountType

import (
	"account-app/internal/model"
	"account-app/internal/model/repository"
	"github.com/devfeel/mapper"
	"log"
)

type accountTypeRepo interface {
	GetUserTypes() (availableTypes []repository.AccountType, err error)
}

type Service struct {
	accountTypeRepo accountTypeRepo
}

func New(repo accountTypeRepo) *Service {
	return &Service{
		accountTypeRepo: repo,
	}
}

func (service *Service) GetForNewUser() (planTypes []model.AccountType) {
	result, err := service.accountTypeRepo.GetUserTypes()
	if err != nil {
		log.Fatalf("Error fetching accountType types %v", err)
	}

	if err = mapper.MapperSlice(&result, &planTypes); err != nil {
		log.Fatalf("Error durign plan type mapping %v", err)
	}

	return planTypes
}
