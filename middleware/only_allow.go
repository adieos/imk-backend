package middleware

import (
	"fmt"
	"net/http"

	"github.com/adieos/imk-backend/dto"
	"github.com/adieos/imk-backend/utils"
	"github.com/gin-gonic/gin"
)

func OnlyAllow(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userRole := ctx.MustGet("role").(string)

		for _, role := range roles {
			if userRole == role {
				ctx.Next()
				return
			}
		}

		err := fmt.Sprintf(dto.ErrRoleNotAllowed.Error(), userRole)
		response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_TOKEN_NOT_VALID, err, nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
	}
}
