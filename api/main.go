package main

import (
	"api/album"
	"api/client"
	"api/transaction"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("REST API - live")
	router := gin.Default()
	router.GET("/album", album.GetAlbums)
	router.POST("/album", album.PostAlbum)
	router.GET("/album/:id", album.GetAlbumByID)

	router.POST("/client", client.PostClient)
	router.GET("/client/:id", client.GetClientByID)

	router.GET("/transaction/:id", transaction.GetTransactionByID)
	router.POST("/transaction", transaction.PostTransaction)

	router.Run(":8080")
}
