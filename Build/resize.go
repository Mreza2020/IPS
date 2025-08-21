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

type ResizeS struct {
	Width  string `json:"width"`
	Height string `json:"height"`
	File   string `json:"file"`
}

func Resize(c *gin.Context) {
	var resize ResizeS
	if err := c.ShouldBindJSON(&resize); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON request"})
	}
	width, err := strconv.Atoi(resize.Width)
	height, err := strconv.Atoi(resize.Height)
	if err != nil || width < 0 || height < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request (width or height)"})
		return
	}
	file, err := os.Open(resize.File)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File not found"})
		return
	}
	defer func(file *os.File) {
		err = file.Close()
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
	dstImage := imaging.Resize(srcImage, width, height, imaging.Lanczos)
	var buf bytes.Buffer
	opts := &jpeg.Options{Quality: 100}
	if err = jpeg.Encode(&buf, dstImage, opts); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode image"})
		return
	}

	c.Data(http.StatusOK, "image/jpeg", buf.Bytes())

}
