package app

import (
	"go-gin-example/pkg/e"

	"github.com/gin-gonic/gin"
)

// Gin gin
type Gin struct {
	C *gin.Context
}

// Response response message
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": httpCode,
		"msg":  e.GetMsg(errCode),
		"data": data,
	})
	return
}
