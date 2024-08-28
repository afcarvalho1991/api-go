package transaction

import (
	"api/album"
	"api/client"
	"encoding/json"
	"fmt"
	"io"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// Transaction represents data about a Transaction.
type Transaction struct {
	ID        uuid.UUID `json:"id"`
	Client    string    `json:"client"`
	Album     string    `json:"album"`
	Amount    float32   `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
}

// albums slice to seed record album data.
var transactions map[uuid.UUID]Transaction = make(map[uuid.UUID]Transaction)

// postClient adds a client from JSON received in the request body.
func PostTransaction(c *gin.Context) {
	var new_tx Transaction

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&new_tx); err != nil {
		return
	}

	// Check if album exist
	response, err := http.Get(fmt.Sprintf("http://localhost:8080/album/" + new_tx.Album))

	if err != nil {
		fmt.Print(err.Error())
		return
	}

	var album album.Album
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	json.Unmarshal(responseData, &album)

	if album.ID == uuid.Nil {
		c.IndentedJSON(http.StatusNotFound, "Album "+new_tx.Album+" not found")
		return
	}

	// Check if Client exist

	response, err = http.Get(fmt.Sprintf("http://localhost:8080/client/" + new_tx.Client))

	if err != nil {
		fmt.Print(err.Error())
		return
	}

	var client client.Client
	responseData, err = io.ReadAll(response.Body)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	json.Unmarshal(responseData, &client)

	if client.ID == uuid.Nil {
		c.IndentedJSON(http.StatusNotFound, "Client "+new_tx.Client+" not found")
		return
	}

	id := createTransaction(new_tx)

	c.IndentedJSON(http.StatusCreated, id)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetTransactionByID(c *gin.Context) {
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

func createTransaction(tx Transaction) uuid.UUID {

	// Add ID & timestamp
	tx.ID = uuid.NewV4()
	tx.Timestamp = time.Now()

	// Add a new transaction to the slice.
	transactions[tx.ID] = tx

	return tx.ID
}
