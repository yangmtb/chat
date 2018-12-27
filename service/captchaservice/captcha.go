package captchaservice

import (
	"chat/pkg/captcha"
	"chat/pkg/setting"
	"fmt"
)

// Captcha struct
type Captcha struct {
	ID    string
	Data  string
	Value string
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
func (c *Captcha) GetCaptcha() {
	id, ins := captcha.GenerateCaptcha("", conf)
	fmt.Println("id:", id, "value:", captcha.GetVerify(id))
	c.ID = id
	c.Data = captcha.WriteToBase64Encoding(ins)
}

// VerifyCaptcha verify captcha
func (c *Captcha) VerifyCaptcha() bool {
	return captcha.VerifyCaptcha(c.ID, c.Value)
}
