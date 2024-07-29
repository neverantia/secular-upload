package secular

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "query received",
	})
}

func Upload(c *gin.Context) {
	file, fileheader, err := c.Request.FormFile("file")

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "error",
		})
		panic(err)
	}

	defer file.Close()

	fileName := fileheader.Filename
	fileExtention := filepath.Ext(fileName)

	uuid := uuid.New()

	filePath := uuid.String() + fileExtention

	path, err := os.Create("./uploads/" + filePath)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "error",
		})
		panic(err)

	}

	defer path.Close()

	_, err = io.Copy(path, file)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "error",
		})
		panic(err)

	}

	c.JSON(http.StatusOK, gin.H{
		"link": filePath,
	})

}
