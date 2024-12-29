package routers

import (
	"trial/controllers"

	"github.com/gin-gonic/gin"
)

func RoutersAuth(rg *gin.RouterGroup) {
	rg.POST("/login", controllers.LoginController)
	rg.POST("/register", controllers.RegisterController)
	rg.POST("/logout", controllers.LogoutController)
}
