package Build

import (
	"bytes"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"image/jpeg"
	"net/http"
	"os"
	"strconv"
)

type CompressST struct {
	File    string `json:"file"`
	Quality string `json:"quality"`
}

func Compress(c *gin.Context) {
	var compress CompressST
	if err := c.ShouldBindJSON(&compress); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON request"})
		return
	}
	file, err := os.Open(compress.File)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File not found"})
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"response": "Failed to properly close the file"})
			return
		}
	}(file)
	srcImage, err := imaging.Decode(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image format"})
		return
	}
	quality := 75
	if val, err1 := strconv.Atoi(compress.Quality); err1 == nil && val >= 1 && val <= 100 {
		quality = val
	}
	var buf bytes.Buffer

	if err = jpeg.Encode(&buf, srcImage, &jpeg.Options{Quality: quality}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode image"})
		return
	}
	c.Data(http.StatusOK, "image/jpeg", buf.Bytes())

}
