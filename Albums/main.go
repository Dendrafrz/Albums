package main

//calling gin, gorm and http
import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:8081)/Albums"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&album{})
	DB = database
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumsByID)
	router.POST("/albums", postAlbums)
	router.Run("localhost:8081")

}

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// record Album Data
var albums = []album{
	{ID: "1", Title: "Keren", Artist: "Dendra", Price: 2000000},

	{ID: "2", Title: "Mantap", Artist: "Naufal", Price: 400000},

	{ID: "3", Title: "ok", Artist: "Fahrezi", Price: 2300000},
}

// GET
// present data to record || gin.Context untuk responds validasi membuat
// StatusOK konsta sebagai index
// "IndentedJSON" panggilan ""Context".JSON"
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// Post
// make variable new Album
// call BindJSON to bind received request Body
// albums as newAlbum created
func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// Get
func getAlbumsByID(c *gin.Context) {
	id := c.Param("id")

	for _, b := range albums {
		if b.ID == id {
			c.IndentedJSON(http.StatusOK, b)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album Not Found"})
}
