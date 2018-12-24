package routers

import (
	"chat/pkg/setting"
	"chat/routers/api"

	"github.com/gin-gonic/gin"
)

// InitRouter is new a gin engine
func InitRouter() (r *gin.Engine) {
	r = gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.ServerSetting.RunMode)

	r.POST("/register", api.Register)
	return
}
