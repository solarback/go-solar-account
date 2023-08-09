package auth

import (
	"account-app/cmd/service/auth"
	"account-app/internal/model/requests"
	"github.com/gin-gonic/gin"
	"net/http"
)

type authService interface {
	Login(login requests.Login) (auth.TokenPairs, error)
}

type Handler struct {
	authService authService
}

func New(service authService) *Handler {
	return &Handler{
		authService: service,
	}
}

func (handler *Handler) Login(c *gin.Context) {
	var loginRequest requests.Login
	if err := c.BindJSON(&loginRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Ooops, something wrong!!!"})
		return
	}

	result, err := handler.authService.Login(loginRequest)

	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Bad credentials"})
	} else {
		c.IndentedJSON(http.StatusOK, result)
	}
}

func (handler *Handler) Logout(c *gin.Context) {

}
