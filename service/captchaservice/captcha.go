package captchaservice

import (
	"chat/pkg/captcha"
	"chat/pkg/setting"
	"fmt"
)

// CaptchaParam struct
type CaptchaParam struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

var (
	conf captcha.ConfigCharacter
)

// SetConfig .
func SetConfig() {
	conf.Height = setting.CaptchaSetting.Height
	conf.Width = setting.CaptchaSetting.Width
	conf.CaptchaLen = setting.CaptchaSetting.Len
	conf.IsUseSimpleFont = setting.CaptchaSetting.UseSimpleFont
	conf.Mode = setting.CaptchaSetting.Mode
	conf.IsShowNoiseDot = setting.CaptchaSetting.DotNoise
	conf.IsShowSlimeLine = setting.CaptchaSetting.SlimeLine
}

// GetCaptcha get captcha
func GetCaptcha() (id, data string) {
	id, ins := captcha.GenerateCaptcha("", conf)
	fmt.Println("id:", id, "value:", captcha.GetVerify(id))
	data = captcha.WriteToBase64Encoding(ins)
	return
}

// VerifyCaptcha verify captcha
func VerifyCaptcha(param CaptchaParam) bool {
	return captcha.VerifyCaptcha(param.ID, param.Value)
}
