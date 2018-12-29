package captcha

import (
	. "chat/pkg/app"
	"chat/pkg/e"
	"chat/service/captchaservice"
	"fmt"

	"github.com/gin-gonic/gin"
)

// Captcha middleware captcha
func Captcha() gin.HandlerFunc {
	return func(c *gin.Context) {
		var param captchaservice.CaptchaParam
		httpCode, code := BindAndValid(c, param)
		if e.SUCCESS != code {
			fmt.Println("not success")
			Response(c, httpCode, code, nil)
		} else if !captchaservice.VerifyCaptcha(param) {
			fmt.Println("not value:", param)
			Response(c, httpCode, e.ERROR_CAPTCHA_VERIFY_FAIL, nil)
		}
		fmt.Println("next")
		c.Next()
	}
}
