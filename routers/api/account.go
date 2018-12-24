package api

import (
	"fmt"
	"log"
	"net/http"

	"chat/service/accountservice"

	"github.com/gin-gonic/gin"
)

// Register to api
func Register(c *gin.Context) {
	var account accountservice.Account
	err := c.Bind(&account)
	if nil != err {
		c.String(http.StatusBadRequest, "param error")
		log.Fatal("bind:", err)
	}
	fmt.Println("account:", account)
	err = account.Register()
	if nil != err {
		c.String(http.StatusInternalServerError, "internal server error")
	} else {
		c.String(http.StatusOK, "ok:", account)
	}
}
