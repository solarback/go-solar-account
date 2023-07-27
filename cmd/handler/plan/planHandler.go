package plan

import (
	"account-app/cmd/domain"
	"github.com/gin-gonic/gin"
)

type service interface {
	GetByName(name string) domain.Plan
}

type Handler struct {
	planService service
}

func (handler *Handler) GetAvailable(ctx *gin.Context) {

}
