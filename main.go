package main

import (
	handlers "pitzdev/web-service-gin/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", handlers.ListAlbums)
	router.GET("/albums/:id", handlers.GetAlbum)
	router.POST("/albums", handlers.CreateAlbum)

	router.Run("localhost:8080")
}
∂