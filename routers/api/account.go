package api

import (
	"chat/pkg/app"
	"chat/pkg/e"
	"chat/pkg/util"
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

// Login to api
func Login(c *gin.Context) {
	appG := app.Gin{C: c}
	var account accountservice.Account
	httpCode, errCode := app.BindAndValid(c, &account)
	if e.SUCCESS != errCode {
		appG.Response(httpCode, errCode, nil)
		return
	}
	if account.Auth() {
		type t struct {
			Token string
		}
		var tt t
		tt.Token, _ = util.GenerateToken(account.Username, account.Password)
		appG.Response(httpCode, errCode, tt)
	} else {
		appG.Response(httpCode, errCode, nil)
	}
}
