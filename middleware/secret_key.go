package middleware

import (
	"net/http"
	"os"

	"github.com/adieos/imk-backend/dto"
	"github.com/adieos/imk-backend/utils"
	"github.com/gin-gonic/gin"
)

func SecretServer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("SERVER_SECRET_KEY")
		if authHeader == "" {
			response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, dto.MESSAGE_FAILED_TOKEN_NOT_FOUND, nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		secretKey := os.Getenv("SERVER_SECRET_KEY")
		if authHeader != secretKey {
			response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, dto.MESSAGE_FAILED_TOKEN_NOT_ALLOWED, nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		ctx.Next()
	}
}
