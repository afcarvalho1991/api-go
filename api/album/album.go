package album

import (
	"api/cpg"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	uuid "github.com/satori/go.uuid"
)

var POSTGRES_CONN string = "postgres://XPTO:XPTO123@postgres-db:5432/album_store?pool_max_conns=10"

// album represents data about a record album.
type Album struct {
	ID     uuid.UUID `json:"id"`
	Title  string    `json:"title"`
	Artist string    `json:"artist"`
	Price  float64   `json:"price"`
}

// getAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
	albums, err := getAllAlbums(c)
	if err != nil {
		fmt.Println("Error - GetAlbums - ", err)
		c.IndentedJSON(
			400,
			gin.H{
				"success": true,
				"message": "Error - GetAlbums - generic error"})
	}

	c.IndentedJSON(
		http.StatusOK,
		gin.H{
			"success": true,
			"message": albums})

}

// postAlbum adds an album from JSON received in the request body.
func PostAlbum(c *gin.Context) {
	var newAlbum Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(
			400,
			gin.H{
				"success": false,
				"message": "invalid album data"})
		return
	}

	err := addAlbum(c, newAlbum)
	if err != nil {
		fmt.Println("Error - PostAlbum - ", err)
		c.IndentedJSON(
			500,
			gin.H{
				"success": false,
				"message": fmt.Sprint("unable to create album --", newAlbum)})
		return
	}

	c.IndentedJSON(http.StatusCreated,
		gin.H{
			"success": true,
			"message": "album created"})
}

// GetAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	album, contains, err := getAlbumByID(c, id)
	if err != nil {
		c.IndentedJSON(
			500,
			gin.H{
				"success": false,
				"message": fmt.Sprint("failed to search by id: ", err)})
		return
	}
	if contains {
		c.IndentedJSON(
			http.StatusOK,
			gin.H{
				"success": true,
				"message": album})
		return

	} else {
		c.IndentedJSON(
			http.StatusNotFound,
			gin.H{
				"success": false,
				"message": "album not found"})
	}
}

// DB interaction

func getAlbumByID(c *gin.Context, id string) (Album, bool, error) {

	// Get DB conn
	pg, err := cpg.NewPG(c, POSTGRES_CONN)
	if err != nil {
		fmt.Println(err)
	}

	// Query
	query := fmt.Sprintf("SELECT ID, Artist, Title, Price FROM album WHERE ID='%s'", id)
	row := pg.DB.QueryRow(c, query)

	album := Album{}
	err = row.Scan(&album.ID, &album.Artist, &album.Title, &album.Price)
	if err != nil {
		return album, false, fmt.Errorf("unable to scan row: %w", err)
	}

	return album, true, nil
}

func getAllAlbums(c *gin.Context) ([]Album, error) {

	// Get DB conn
	pg, err := cpg.NewPG(c, POSTGRES_CONN)
	if err != nil {
		fmt.Println(err)
	}

	// Query
	query := `SELECT ID, Artist, Title, Price FROM album`

	rows, err := pg.DB.Query(c, query)
	if err != nil {
		return nil, fmt.Errorf("unable to query albums: %w", err)
	}

	defer rows.Close()

	albums := []Album{}
	for rows.Next() {
		album := Album{}
		err := rows.Scan(&album.ID, &album.Artist, &album.Title, &album.Price)
		if err != nil {
			return nil, fmt.Errorf("unable to scan row: %w", err)
		}
		albums = append(albums, album)
	}

	return albums, nil
}

func addAlbum(ctx *gin.Context, newAlbum Album) error {

	pg, err := cpg.NewPG(ctx, POSTGRES_CONN)
	if err != nil {
		fmt.Println(err)
	}

	query := `INSERT INTO Album (Title, Artist, Price) VALUES (@Title, @Artist, @Price) RETURNING ID;`

	args := pgx.NamedArgs{
		"Title":  newAlbum.Title,
		"Artist": newAlbum.Artist,
		"Price":  newAlbum.Price,
	}

	_, err = pg.DB.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}
	return nil
}
