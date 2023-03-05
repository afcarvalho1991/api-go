package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/album", getAlbums)
	router.POST("/album", postAlbums)
	router.GET("/album/:id", getAlbumByID)

	router.POST("/client", postClient)
	router.GET("/client/:id", getClientByID)

	router.GET("/transaction/:id", getTransactionByID)
	router.POST("/transaction", postTransaction)

	router.Run("localhost:8080")
}
