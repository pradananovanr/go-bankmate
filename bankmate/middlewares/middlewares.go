package middlewares

import (
	"go-paybro/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, err := util.TokenValid(ctx)
		if err != nil {
			ctx.String(http.StatusUnauthorized, "unauthorized")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
