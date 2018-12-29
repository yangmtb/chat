package api

import (
	. "chat/pkg/app"
	"chat/pkg/e"
	"chat/pkg/util"
	"chat/service/accountservice"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Signup to api
func Signup(c *gin.Context) {
	var account accountservice.Account
	httpCode, errCode := BindAndValid(c, &account.Params)
	if e.SUCCESS != errCode {
		Response(c, httpCode, errCode, nil)
		return
	}
	fmt.Println("account:", account)
	err := account.Signup()
	if nil != err {
		Response(c, http.StatusInternalServerError, e.ERROR_ACCOUNT_SIGN_UP_FAIL, nil)
	} else {
		Response(c, httpCode, e.SUCCESS, nil)
	}
}

// Signin to api
func Signin(c *gin.Context) {
	var account accountservice.Account
	httpCode, errCode := BindAndValid(c, &account.Params)
	if e.SUCCESS != errCode {
		Response(c, httpCode, errCode, nil)
		return
	}
	fmt.Println("signin:", account.Params)
	if account.Auth() {
		type t struct {
			Token string
		}
		var tt t
		tt.Token, _ = util.GenerateToken(account.Params.Username, account.Params.Password)
		//token, _ := c.Cookie("token")
		c.SetCookie("token", tt.Token, 300, "/", "localhost", false, true)
		Response(c, httpCode, e.SUCCESS, tt)
	} else {
		Response(c, httpCode, e.ERROR_ACCOUNT_SIGN_IN_FAIL, nil)
	}
}

// Exist to api phone,email,username
func Exist(c *gin.Context) {
	var account accountservice.Account
	type t struct {
		Key   string
		Value string
	}
	var tt t
	httpCode, errCode := BindAndValid(c, &tt)
	if e.SUCCESS != errCode {
		Response(c, httpCode, errCode, nil)
		return
	}
	if account.Exist(tt.Key, tt.Value) {
		if "phone" == tt.Key {
			errCode = e.ERROR_ACCOUNT_PHONE_EXIST
		} else if "email" == tt.Key {
			errCode = e.ERROR_ACCOUNT_EMAIL_EXIST
		} else if "username" == tt.Key {
			errCode = e.ERROR_ACCOUNT_USERNAME_EXIST
		} else {
			errCode = e.ERROR
		}
		Response(c, httpCode, errCode, nil)
	} else {
		Response(c, httpCode, e.SUCCESS, nil)
	}
}
