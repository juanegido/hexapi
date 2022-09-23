package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckHealth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "OK")
	}
}
