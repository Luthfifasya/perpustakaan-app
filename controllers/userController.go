package controllers

import (
	"net/http"
	"perpustakaan-app/utils"

	"github.com/gin-gonic/gin"
)

// Login dan generate JWT
func UserLogin(c *gin.Context) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Validasi username & password
	if creds.Username != "admin" || creds.Password != "password123" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username atau Password salah"})
		return
	}

	// generate JWT
	token, err := utils.GenerateJWT(creds.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Tidak bisa mengenerate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"username": creds.Username,
		"token":    token,
	})
}
