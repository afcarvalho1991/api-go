package album

import (
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// album represents data about a record album.
type Album struct {
	ID     uuid.UUID `json:"id"`
	Title  string    `json:"title"`
	Artist string    `json:"artist"`
	Price  float64   `json:"price"`
}

// albums slice to seed record album data.
var albums map[uuid.UUID]Album = make(map[uuid.UUID]Album)

// getAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, getAllAlbums())
}

// postAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context) {
	var newAlbum Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	newAlbum = addAlbum(newAlbum)

	c.IndentedJSON(http.StatusCreated, newAlbum.ID)
}

// GetAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	album, contains := getAlbumByID(id)

	if contains {
		c.IndentedJSON(http.StatusOK, album)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// DB interaction

func getAlbumByID(id string) (Album, bool) {
	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	album, contains := albums[uuid.FromStringOrNil(id)]
	return album, contains
}

func addAlbum(newAlbum Album) Album {

	// Add ID
	newAlbum.ID = uuid.NewV4()

	// Add the new album to the slice.
	albums[newAlbum.ID] = newAlbum

	return newAlbum
}

func getAllAlbums() map[uuid.UUID]Album {
	return albums
}
