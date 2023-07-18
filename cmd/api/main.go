package main

import (
	"account-app/cmd/account"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

const port = 8080

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/accounts", account.GetAvailableAccounts)
	router.POST("/addAccount", account.AddNewAccount)

	router.Run(fmt.Sprintf(":%d", port))
}
