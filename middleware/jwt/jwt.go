package jwt

import (
	. "chat/pkg/app"
	"chat/pkg/e"
	"chat/pkg/util"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWT .
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := e.SUCCESS
		token := c.Query("token")
		if "" == token {
			code = e.INVALID_PARAMS
		} else {
			_, err := util.ParseToken(token)
			if nil != err {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.ERROR_TOKEN_TIMEOUT_FAIL
				default:
					code = e.ERROR_TOKEN_CHECK_FAIL
				}
			}
		}
		if e.SUCCESS != code {
			Response(c, http.StatusNonAuthoritativeInfo, code, nil)
		}
		c.Next()
	}
}
