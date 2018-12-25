package routers

import (
	"chat/pkg/setting"
	"chat/routers/api"
	"chat/routers/api/v1"

	"github.com/gin-gonic/gin"
)

// InitRouter is new a gin engine
func InitRouter() (r *gin.Engine) {
	r = gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.ServerSetting.RunMode)

	r.POST("/register", api.Register)
	r.POST("/login", api.Login)

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/ping", v1.Ping)
	}
	return
}
