package routes

import (
	"github.com/adieos/imk-backend/controller"
	"github.com/adieos/imk-backend/middleware"
	"github.com/adieos/imk-backend/service"
	"github.com/gin-gonic/gin"
)

func User(route *gin.Engine, userController controller.UserController, jwtService service.JWTService) {
	routes := route.Group("/api/auth")
	{
		// User
		routes.POST("", userController.Register)
		routes.POST("/login", userController.Login)
		routes.PUT("/update", middleware.Authenticate(jwtService), userController.Update)
		routes.POST("/sendmail", userController.SendVerificationEmail)
		routes.GET("/verify", userController.VerifyEmail)
		routes.GET("/me", middleware.Authenticate(jwtService), userController.Me)
		routes.POST("/reset", userController.ResetPassword)
		routes.POST("/forget", userController.ForgetPassword)
	}
}
