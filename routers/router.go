package routers

import (
	"chat/middleware/captcha"
	"chat/middleware/jwt"
	"chat/pkg/setting"
	"chat/routers/api"
	v1 "chat/routers/api/v1"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// InitRouter is new a gin engine
func InitRouter() (r *gin.Engine) {
	r = gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.ServerSetting.RunMode)
	//r.Use(static.Serve("/", static.LocalFile("/tmp", false)))
	//r.Use(static.Serve("/", static.LocalFile("./static", false)))
	r.Use(static.Serve("/", static.LocalFile("./routers/static", true)))
	//r.StaticFile("/", "./static")

	r.GET("/captcha", api.GetCaptcha)
	r.POST("/exist", api.Exist)
	r.POST("/verify", api.VerifyCaptcha)

	// 需要验证码的api
	c := r.Group("/api")
	c.Use(captcha.Captcha())
	{
		c.POST("/register", api.Signup)
		c.POST("/login", api.Signin)
	}

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/ping", v1.Ping)
	}
	return
}
