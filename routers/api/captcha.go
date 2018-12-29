package api

import (
	"chat/pkg/app"
	"chat/pkg/e"
	"chat/service/captchaservice"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appG app.Gin
)

// GetCaptcha api
func GetCaptcha(c *gin.Context) {
	appG.C = c
	id, data := captchaservice.GetCaptcha()
	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{"id": id, "data": data})
}

// VerifyCaptcha api
func VerifyCaptcha(c *gin.Context) {
	appG.C = c
	var param captchaservice.CaptchaParam
	httpCode, errCode := appG.BindAndValid(&param)
	if e.SUCCESS != errCode {
		appG.Response(httpCode, errCode, nil)
		return
	}
	if captchaservice.VerifyCaptcha(param) {
		appG.Response(http.StatusOK, e.SUCCESS, nil)
	} else {
		appG.Response(http.StatusOK, e.ERROR_CAPTCHA_VERIFY_FAIL, nil)
	}
}
