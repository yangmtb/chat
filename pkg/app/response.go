package app

import (
	"chat/pkg/e"

	"github.com/gin-gonic/gin"
)

// Response response message
func Response(c *gin.Context, httpCode, errCode int, data interface{}) {
	c.JSON(httpCode, gin.H{
		"code": errCode,
		"msg":  e.GetMsg(errCode),
		"data": data,
	})
	return
}
