package api

import (
	. "chat/pkg/app"
	"chat/pkg/e"
	"chat/service/captchaservice"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCaptcha api
func GetCaptcha(c *gin.Context) {
	id, data := captchaservice.GetCaptcha()
	Response(c, http.StatusOK, e.SUCCESS, map[string]string{"id": id, "data": data})
}

// VerifyCaptcha api
func VerifyCaptcha(c *gin.Context) {
	var param captchaservice.CaptchaParam
	httpCode, errCode := BindAndValid(c, &param)
	if e.SUCCESS != errCode {
		Response(c, httpCode, errCode, nil)
		return
	}
	if captchaservice.VerifyCaptcha(param) {
		Response(c, http.StatusOK, e.SUCCESS, nil)
	} else {
		Response(c, http.StatusOK, e.ERROR_CAPTCHA_VERIFY_FAIL, nil)
	}
}
