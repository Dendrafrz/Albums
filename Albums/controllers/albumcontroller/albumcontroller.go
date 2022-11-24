package albumcontroller

import (
	"net/http"

	"github.com/Dendrafrz/Albums/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	var newAlbum []models.album

	models.DB.Find(&album)
	c.JSON(http.StatusOK, gin.H{"albums": albums})
}

func Create(c *gin.Context) {
	var newAlbum models.album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func Show(c *gin.Context) {
	id := c.Param("id")

	for _, b := range albums {
		if b.ID == id {
			c.IndentedJSON(http.StatusOK, b)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album Not Found"})
}
