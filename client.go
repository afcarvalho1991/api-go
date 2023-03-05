package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// Client represents data about a Client.
type client struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Age  int8      `json:"age"`
}

// clients db
var clients map[uuid.UUID]client = make(map[uuid.UUID]client)

// postClient adds a client from JSON received in the request body.
func postClient(c *gin.Context) {
	var new_client client

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&new_client); err != nil {
		return
	}

	// Add ID
	new_client.ID = uuid.NewV4()

	// Add a new client to the slice.
	clients[new_client.ID] = new_client
	c.IndentedJSON(http.StatusCreated, new_client.ID)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getClientByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	client, contains := clients[uuid.FromStringOrNil(id)]
	if contains {
		c.IndentedJSON(http.StatusOK, client)
		return
	}
	c.IndentedJSON(
		http.StatusNotFound,
		gin.H{"message": "client with id=" + id + " not found"})
}
