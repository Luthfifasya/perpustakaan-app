package controllers

import (
	"net/http"
	"perpustakaan-app/database"
	"perpustakaan-app/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GET /api/books
func GetBooks(c *gin.Context) {
	var books []models.Book
	database.DB.Preload("Category").Find(&books)
	c.JSON(http.StatusOK, books)
}

// POST /api/books
func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi release_year
	if book.ReleaseYear < 1980 || book.ReleaseYear > 2024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "release_year harus antara 1980 - 2024"})
		return
	}

	// Konversi thickness berdasarkan total_page
	if book.TotalPage > 100 {
		book.Thickness = "tebal"
	} else {
		book.Thickness = "tipis"
	}

	if err := database.DB.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	database.DB.Preload("Category").First(&book, book.ID)
	c.JSON(http.StatusCreated, book)
}

// GET /api/books/:id
func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := database.DB.Preload("Category").First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Buku tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, book)
}

// PUT /api/books/:id
func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := database.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Buku tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi release_year
	if book.ReleaseYear < 1980 || book.ReleaseYear > 2024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "release_year harus antara 1980 - 2024"})
		return
	}

	// Konversi thickness
	if book.TotalPage > 100 {
		book.Thickness = "tebal"
	} else {
		book.Thickness = "tipis"
	}

	database.DB.Save(&book)
	c.JSON(http.StatusOK, book)
}

// DELETE /api/books/:id
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	bookID, _ := strconv.Atoi(id)

	// cek apakah buku ada
	var book models.Book
	if err := database.DB.First(&book, bookID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Buku tidak ditemukan"})
		return
	}

	database.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"message": "Buku berhasil dihapus"})
}
