package apirest

import (
	"api/album"
	"api/client"
	"api/transaction"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/album", album.GetAlbums)
	router.POST("/album", album.PostAlbums)
	router.GET("/album/:id", album.GetAlbumByID)

	router.POST("/client", client.PostClient)
	router.GET("/client/:id", client.GetClientByID)

	router.GET("/transaction/:id", transaction.GetTransactionByID)
	router.POST("/transaction", transaction.PostTransaction)

	router.Run("localhost:8080")
}
