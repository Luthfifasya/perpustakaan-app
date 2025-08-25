package routes

import (
	"perpustakaan-app/controllers"
	"perpustakaan-app/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	// Endpoint login user → generate JWT
	r.POST("/api/users/register", controllers.RegisterUser)
	r.GET("/api/users", controllers.GetUsers)
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

	// Books
	books := api.Group("/books")
	books.Use(middleware.JWTAuth())
	{
		books.GET("", controllers.GetBooks)
		books.POST("", controllers.CreateBook)
		books.GET("/:id", controllers.GetBookByID)
		books.DELETE("/:id", controllers.DeleteBook)
	}

}
