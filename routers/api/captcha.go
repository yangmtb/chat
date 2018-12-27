package api

import (
	"chat/pkg/app"
	"chat/pkg/e"
	"chat/service/captchaservice"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCaptcha api
func GetCaptcha(c *gin.Context) {
	appG := app.Gin{C: c}
	var captcha captchaservice.Captcha
	captcha.GetCaptcha()
	appG.Response(http.StatusOK, e.SUCCESS, captcha)
}

// VerifyCaptcha api
func VerifyCaptcha(c *gin.Context) {
	appG := app.Gin{C: c}
	var captcha captchaservice.Captcha
	httpCode, errCode := appG.BindAndValid(&captcha)
	if e.SUCCESS != errCode {
		appG.Response(httpCode, errCode, nil)
		return
	}
	if captcha.VerifyCaptcha() {
		appG.Response(http.StatusOK, e.SUCCESS, nil)
	} else {
		appG.Response(http.StatusOK, e.ERROR_CAPTCHA_VERIFY_FAIL, nil)
	}
}
