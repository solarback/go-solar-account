package account

import (
	"account-app/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAvailableAccounts(c *gin.Context) {

	availableAccounts := getAll()
	c.IndentedJSON(http.StatusOK, *availableAccounts)
}

func AddNewAccount(c *gin.Context) {
	var newAccount model.Account

	if err := c.BindJSON(&newAccount); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Ooops, something wrong!!!"})
		return
	}

	addAccount(&newAccount)
	c.IndentedJSON(http.StatusCreated, newAccount)
}
