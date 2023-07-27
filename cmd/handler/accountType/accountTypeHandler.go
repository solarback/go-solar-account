package accountType

import (
	"account-app/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type service interface {
	GetForNewUser() []model.AccountType
}

type Handler struct {
	accountTypeService service
}

func Init(service service) *Handler {
	return &Handler{
		accountTypeService: service,
	}
}

func (handler *Handler) GetForNewUser(ctx *gin.Context) {
	accountTypeService := handler.accountTypeService
	result := accountTypeService.GetForNewUser()

	ctx.IndentedJSON(http.StatusOK, result)
}
