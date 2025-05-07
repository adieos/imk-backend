package middleware

import (
	"net/http"
	"os"

	"github.com/adieos/imk-backend/dto"
	"github.com/adieos/imk-backend/utils"
	"github.com/gin-gonic/gin"
)

func Xendit() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("X-CALLBACK-TOKEN")
		if authHeader == "" {
			response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, dto.MESSAGE_FAILED_TOKEN_NOT_FOUND, nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		xenditKey := os.Getenv("XENDIT_WEBHOOK_TOKEN")
		if authHeader != xenditKey {
			response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, dto.MESSAGE_FAILED_TOKEN_NOT_ALLOWED, nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		ctx.Next()
	}
}
