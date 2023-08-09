package auth

import (
	"account-app/cmd/domain"
	"account-app/internal/model/requests"
	"errors"
	"fmt"
)

type tokenService interface {
	GenerateTokenPair(user *JwtUser) (TokenPairs, error)
}

type userService interface {
	GetByEmail(email string) (domain.User, error)
}

type Service struct {
	userService  userService
	tokenService tokenService
}

func New(userService userService, tokenService tokenService) *Service {
	return &Service{
		userService:  userService,
		tokenService: tokenService,
	}
}

func (s Service) Login(login requests.Login) (TokenPairs, error) {

	user, err := s.userService.GetByEmail(login.Email)

	if err != nil {
		panic(fmt.Sprintf("User not registered: %v", err))
	}

	if !user.CheckPassword(login.Password) {
		return TokenPairs{}, errors.New("error on auth, password is incorrect")
	}

	return s.tokenService.GenerateTokenPair(
		&JwtUser{
			ID:       user.Id.String(),
			UserName: user.UserName,
		})
}
