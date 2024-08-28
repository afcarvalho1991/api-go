package client

import (
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// Client represents data about a Client.
type Client struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Age  int8      `json:"age"`
}

// clients db
var clients map[uuid.UUID]Client = make(map[uuid.UUID]Client)

// postClient adds a client from JSON received in the request body.
func PostClient(c *gin.Context) {
	var newClient Client

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newClient); err != nil {
		return
	}

	newClient = addClient(newClient)

	c.IndentedJSON(http.StatusCreated, newClient.ID)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetClientByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of Clients, looking for
	// an album whose ID value matches the parameter.
	client, contains := getClientByID(id)

	if contains {
		c.IndentedJSON(http.StatusOK, client)
		return
	}
	c.IndentedJSON(
		http.StatusNotFound,
		gin.H{"message": "client with id=" + id + " not found"})
}

// Comm with DB

func getClientByID(id string) (Client, bool) {
	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	client, contains := clients[uuid.FromStringOrNil(id)]
	return client, contains
}

func addClient(newClient Client) Client {

	// Add ID
	newClient.ID = uuid.NewV4()

	// Add the new Client to the slice.
	clients[newClient.ID] = newClient

	return newClient
}
