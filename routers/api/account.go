package api

import (
	"chat/pkg/app"
	"chat/pkg/e"
	"chat/pkg/util"
	"fmt"
	"net/http"

	"chat/service/accountservice"

	"github.com/gin-gonic/gin"
)

// Signup to api
func Signup(c *gin.Context) {
	appG := app.Gin{C: c}
	var account accountservice.Account
	httpCode, errCode := appG.BindAndValid(&account.Params)
	if e.SUCCESS != errCode {
		appG.Response(httpCode, errCode, nil)
		return
	}
	fmt.Println("account:", account)
	err := account.Signup()
	if nil != err {
		appG.Response(http.StatusInternalServerError, e.ERROR_ACCOUNT_SIGN_UP_FAIL, nil)
	} else {
		appG.Response(httpCode, e.SUCCESS, nil)
	}
}

// Signin to api
func Signin(c *gin.Context) {
	appG := app.Gin{C: c}
	var account accountservice.Account
	httpCode, errCode := appG.BindAndValid(&account.Params)
	if e.SUCCESS != errCode {
		appG.Response(httpCode, errCode, nil)
		return
	}
	if account.Auth() {
		type t struct {
			Token string
		}
		var tt t
		tt.Token, _ = util.GenerateToken(account.Params.Username, account.Params.Password)
		//token, _ := c.Cookie("token")
		c.SetCookie("token", tt.Token, 300, "/", "localhost", false, true)
		appG.Response(httpCode, e.SUCCESS, tt)
	} else {
		appG.Response(httpCode, e.ERROR_ACCOUNT_SIGN_IN_FAIL, nil)
	}
}
