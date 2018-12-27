package main

import (
	"chat/model"
	"chat/pkg/captcha"
	"chat/pkg/logging"
	"chat/pkg/util"
	"chat/routers"
	"chat/service/captchaservice"
	"fmt"
	"log"
	"net/http"
	"os"

	"chat/pkg/setting"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func test() {
	r := gin.Default()
	// if Allow DirectoryIndex
	//r.Use(static.Serve("/", static.LocalFile("/tmp", true)))
	// set prefix
	//r.Use(static.Serve("/static", static.LocalFile("/tmp", true)))
	r.Use(static.Serve("/", static.LocalFile("./routers/static", false)))
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "test")
	})
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8000")
	return
	var conf captcha.ConfigCharacter
	conf.Height = 60
	conf.Width = 240
	conf.CaptchaLen = 7
	conf.IsUseSimpleFont = true
	conf.Mode = 3
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
	psd := `pwd`
	sa := util.GenerateSalt(32)
	fmt.Println("salt:", sa)
	fmt.Println("password:", util.GeneratePassword(psd, sa))
	fmt.Println(util.HashAndSalt([]byte("xxxxxx")))
	return
}

func main() {
	//test()
	//return
	setting.Setup()
	model.Setup()
	logging.Setup()
	captchaservice.SetConfig()

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
