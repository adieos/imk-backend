package middleware

import (
	"os"
	"strings"

	"github.com/adieos/imk-backend/utils/logger"
	"github.com/gin-gonic/gin"
)

func Mayar() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			logger.Errorln("auth header not found")
			ctx.Set("authHeader", 0)
			ctx.Next()
		}
		authHeader = strings.Replace(authHeader, "Bearer ", "", -1)
		xenditKey := os.Getenv("MAYAR_API_KEY")
		if authHeader != xenditKey {
			ctx.Set("authHeader", 0)
			ctx.Next()
		}

		ctx.Set("authHeader", 1)
		ctx.Next()
	}
}
