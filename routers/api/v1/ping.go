package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping pong
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
