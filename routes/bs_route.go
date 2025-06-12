package routes

import (
	"github.com/adieos/imk-backend/controller"
	"github.com/adieos/imk-backend/middleware"
	"github.com/adieos/imk-backend/service"
	"github.com/gin-gonic/gin"
)

func BS(route *gin.Engine, bsController controller.BSController, jwtService service.JWTService) {
	routes := route.Group("/api/bs")
	{
		// Bank Sampah
		routes.POST("", middleware.Authenticate(jwtService), bsController.CreateBS)
		routes.GET("/:id", bsController.GetBSById)
		routes.GET("", bsController.GetAllBS)
		routes.GET("/my-bank-sampah", middleware.Authenticate(jwtService), bsController.GetAllBSByUserId)
		routes.PUT("/update", middleware.Authenticate(jwtService), bsController.UpdateBS)
		routes.PATCH("/:id/status", middleware.Authenticate(jwtService), bsController.ChangeStatusBS)
	}
}
