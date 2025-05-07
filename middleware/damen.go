package middleware

import (
	"net/http"

	"github.com/adieos/imk-backend/constants"
	"github.com/adieos/imk-backend/dto"
	"github.com/adieos/imk-backend/utils"
	"github.com/gin-gonic/gin"
)

func Damen() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid := ctx.MustGet("user_id")
		if uid != constants.DAMEN_UID {
			response := utils.BuildResponseFailed(dto.PESAN_AKSI_TIDAK_DIPERBOLEHKAN, dto.MESSAGE_FAILED_PROSES_REQUEST, nil)
			ctx.AbortWithStatusJSON(http.StatusForbidden, response)
			return
		}

		ctx.Next()
	}
}
