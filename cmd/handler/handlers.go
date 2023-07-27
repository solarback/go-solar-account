package handler

import (
	"account-app/cmd/handler/accountType"
	"account-app/cmd/handler/plan"
	"account-app/cmd/handler/user"
	"account-app/cmd/service"
	"account-app/internal/initializers"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type Handlers struct {
	UserHandler        *user.Handler
	PlanHandler        *plan.Handler
	AccountTypeHandler *accountType.Handler
}

func Init(services *service.Services) *Handlers {
	return &Handlers{
		UserHandler:        user.Init(services.UserService),
		AccountTypeHandler: accountType.Init(services.AccountTypeService),
	}
}

func (handlers *Handlers) RegisterRoutes(config initializers.EnvironmentConfig) {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}),
	)

	router.GET("/users", handlers.UserHandler.GetAllUsers)
	router.POST("/addAccount", handlers.UserHandler.RegisterNewUser)

	router.GET("/availablePlans", handlers.PlanHandler.GetAvailable)

	router.GET("/userAccountTypes", handlers.AccountTypeHandler.GetForNewUser)

	err := router.Run(fmt.Sprintf(":%d", config.ServerPort))
	if err != nil {
		log.Fatalf("Error on registering handlers %v", err)
	}
}
