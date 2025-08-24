package routes

import (
	"perpustakaan-app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Endpoint login user → generate JWT
	r.POST("/api/users/login", controllers.UserLogin)

}
