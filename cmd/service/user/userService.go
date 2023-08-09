package user

import (
	. "account-app/cmd/domain"
	"account-app/internal/model/requests"
	"github.com/devfeel/mapper"
	"github.com/google/uuid"
	"log"
)

type userRepo interface {
	AddUser(newUser User) (User, error)
	GetById(id uuid.UUID) (User, error)
	GetAll() ([]User, error)
}

type Service struct {
	repository userRepo
}

func New(userRepo userRepo) *Service {
	return &Service{
		repository: userRepo,
	}
}

func (service *Service) RegisterNewUser(request requests.RegisterUser) User {
	var err error
	newUser := Create(request.UserName, request.Password)
	newUser.AddEmail(request.Email)
	newUser.AssignAccount(request.AccountType, "FREE")

	if err != nil {
		log.Fatalf("Error on fetching plan by name=%s %v", "Free", err)
	}

	result, err := service.repository.AddUser(newUser)

	if err != nil {
		log.Fatalf("Error on saving new user %v, error=%v", newUser, err)
	}

	return result
}

func (service *Service) GetAll() []User {
	var result []User
	users, err := service.repository.GetAll()

	if err != nil {
		log.Fatalf("Error raeding users from DB %v", err)
	}

	if err = mapper.MapperSlice(users, result); err != nil {
		log.Fatalf("Error on mapping %v", err)
	}

	return result
}

func (service *Service) GetById(id string) User {
	userId, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Error during parsing planId %v", err)
	}

	result, err := service.repository.GetById(userId)
	if err != nil {
		log.Fatalf("Error during fetching user by id=%s %v", id, err)
	}

	return result
}
