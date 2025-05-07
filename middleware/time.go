package middleware

import (
	"net/http"
	"time"

	"github.com/adieos/imk-backend/dto"
	"github.com/adieos/imk-backend/utils"
	"github.com/gin-gonic/gin"
)

// Reject if request is made before limit. FORMAT: YYYY-MM-DD hh:mm:ss
func NotBefore(limit string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		waktu, err := time.Parse("2006-01-02 15:04:05", limit)
		if err != nil {
			response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, dto.MESSAGE_FAILED_PARSE_TIME, nil)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
			return
		}

		now := time.Now()
		if now.Before(waktu) {
			response := utils.BuildResponseFailed(dto.PESAN_DILUAR_MASA_REGISTRASI, dto.MESSAGE_FAILED_PROSES_REQUEST, nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		ctx.Next()
	}
}

// Reject if request is made after limit. FORMAT: YYYY-MM-DD hh:mm:ss
func NotAfter(limit string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		waktu, err := time.Parse("2006-01-02 15:04:05", limit)
		if err != nil {
			response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, dto.MESSAGE_FAILED_PARSE_TIME, nil)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
			return
		}

		now := time.Now()
		if now.After(waktu) {
			response := utils.BuildResponseFailed(dto.PESAN_DILUAR_MASA_REGISTRASI, dto.PESAN_DILUAR_MASA_REGISTRASI, nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		ctx.Next()
	}
}
