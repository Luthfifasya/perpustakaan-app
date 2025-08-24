package routes

import (
	"perpustakaan-app/controllers"
	"perpustakaan-app/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	// Endpoint login user → generate JWT
	r.POST("/api/users/login", controllers.UserLogin)

	// Categories
	categories := api.Group("/categories")
	categories.Use(middleware.JWTAuth())
	{
		categories.GET("", controllers.GetCategories)
		categories.POST("", controllers.CreateCategory)
		categories.GET("/:id", controllers.GetCategoryByID)
		categories.DELETE("/:id", controllers.DeleteCategory)
		categories.GET("/:id/books", controllers.GetBooksByCategory)
	}
}
