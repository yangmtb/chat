package main

import (
	"chat/model"
	"chat/pkg/util"
	"chat/routers"
	"fmt"
	"log"
	"net/http"

	"chat/pkg/setting"
)

func main() {
	psd := `pwd`
	fmt.Println("psd:", psd)
	s1 := util.Sha256(psd)
	fmt.Println("sha256:", s1)
	sa := util.GenerateSalt(32)
	fmt.Println("salt:", sa)
	fmt.Println(util.Sha256(s1 + sa))
	fmt.Println("len:", len(s1))
	fmt.Println(util.HashAndSalt([]byte("xxxxxx")))

	return
	setting.Setup()
	model.Setup()
	fmt.Println("app", setting.AppSetting)
	fmt.Println("server", setting.ServerSetting)
	fmt.Println("database", setting.DatabaseSetting)
	fmt.Println("redis", setting.RedisSetting)
	router := routers.InitRouter()
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("[info] start http server listening on %d", setting.ServerSetting.HTTPPort)
	server.ListenAndServe()
}
