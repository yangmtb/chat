package main

import (
	"chat/model"
	"chat/pkg/captcha"
	"chat/pkg/logging"
	"chat/pkg/util"
	"chat/routers"
	"fmt"
	"log"
	"net/http"
	"os"

	"chat/pkg/setting"
)

func test() {
	psd := `pwd`
	sa := util.GenerateSalt(32)
	fmt.Println("salt:", sa)
	fmt.Println("password:", util.GeneratePassword(psd, sa))
	fmt.Println(util.HashAndSalt([]byte("xxxxxx")))
	return
}

func main() {
	var conf captcha.ConfigCharacter
	conf.Height = 60
	conf.Width = 240
	conf.CaptchaLen = 7
	conf.IsUseSimpleFont = true
	conf.Mode = 2
	id, ins := captcha.GenerateCaptcha("", conf)
	f, err := os.OpenFile("xx.png", os.O_RDWR|os.O_CREATE, 0755)
	if nil != err {
		fmt.Println(err)
		return
	}
	n, err := ins.WriteTo(f)
	fmt.Println(id, n, err)
	fmt.Println(captcha.GetVerify(id))
	return
	setting.Setup()
	model.Setup()
	logging.Setup()

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HTTPPort),
		Handler:        routers.InitRouter(),
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("[info] start http server listening on %d", setting.ServerSetting.HTTPPort)
	server.ListenAndServe()
}
