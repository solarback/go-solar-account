package main

import (
	"account-app/cmd/account"
	"fmt"
	"github.com/gin-gonic/gin"
)

const port = 8080

func main() {
	router := gin.Default()

	router.GET("/accounts", account.GetAvailableAccounts)
	router.POST("/addAccount", account.AddNewAccount)

	router.Run(fmt.Sprintf(":%d", port))
}
