package app

import (
	"chat/pkg/e"

	"github.com/gin-gonic/gin"
)

// Gin gin
type Gin struct {
	C *gin.Context
}

// Response response message
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": errCode,
		"msg":  e.GetMsg(errCode),
		"data": data,
	})
	return
}
