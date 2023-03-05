package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// Transaction represents data about a Transaction.
type transaction struct {
	ID        uuid.UUID `json:"id"`
	Client    string    `json:"client"`
	Album     string    `json:"album"`
	Amount    float32   `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
}

// albums slice to seed record album data.
var transactions map[uuid.UUID]transaction = make(map[uuid.UUID]transaction)

// postClient adds a client from JSON received in the request body.
func postTransaction(c *gin.Context) {
	var new_tx transaction

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&new_tx); err != nil {
		return
	}

	// Check if album exist
	_, hasAlbum := albums[uuid.FromStringOrNil(new_tx.Album)]
	if !hasAlbum {
		c.IndentedJSON(http.StatusNotFound, "Album "+new_tx.Album+" not found")
		return
	}

	// Check if client exist
	_, hasClient := clients[uuid.FromStringOrNil(new_tx.Client)]
	if !hasClient {
		c.IndentedJSON(http.StatusNotFound, "Client "+new_tx.Client+" not found")
		return
	}

	// Add ID
	new_tx.ID = uuid.NewV4()
	new_tx.Timestamp = time.Now()

	// Add a new client to the slice.
	transactions[new_tx.ID] = new_tx
	c.IndentedJSON(http.StatusCreated, new_tx.ID)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getTransactionByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	transaction, contains := transactions[uuid.FromStringOrNil(id)]
	// If the key exists
	if contains {
		c.IndentedJSON(http.StatusOK, transaction)
		return
	}

	// Not found
	c.IndentedJSON(
		http.StatusNotFound,
		gin.H{"message": "transaction with id=" + id + " not found"})
}
