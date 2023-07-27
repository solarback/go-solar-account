package user

import (
	"account-app/cmd/domain"
	"account-app/cmd/service/user"
	"account-app/internal/model/requests"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userService interface {
	GetAll() []domain.User
	GetById(id string) domain.User
	RegisterNewUser(user requests.RegisterUser) domain.User
}

type Handler struct {
	userService userService
}

func Init(service *user.Service) *Handler {
	return &Handler{
		userService: service,
	}
}

func (handler *Handler) GetAllUsers(c *gin.Context) {
	userService := handler.userService
	availableAccounts := userService.GetAll()
	c.IndentedJSON(http.StatusOK, availableAccounts)
}

func (handler *Handler) RegisterNewUser(c *gin.Context) {
	userService := handler.userService
	var newAccount requests.RegisterUser

	if err := c.BindJSON(&newAccount); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Ooops, something wrong!!!"})
		return
	}

	userService.RegisterNewUser(newAccount)
	c.IndentedJSON(http.StatusCreated, newAccount)
}
