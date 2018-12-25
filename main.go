package main

import (
	"chat/model"
	"chat/pkg/logging"
	"chat/pkg/util"
	"chat/routers"
	"fmt"
	"log"
	"net/http"

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
