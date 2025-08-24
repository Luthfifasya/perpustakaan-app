package main

import (
	"fmt"
	"log"
	"os"

	"perpustakaan-app/database"
	"perpustakaan-app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	//inisialisasi database
	database.InitDB()

	//setup router
	r := gin.Default()
	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := fmt.Sprintf(":%s", port)
	fmt.Println("Server running on http://localhost" + addr)
	if err := r.Run(addr); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
