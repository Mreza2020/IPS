package Build

import (
	"bytes"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"image"
	"image/jpeg"
	"net/http"
	"os"
)

type FlipST struct {
	File string `json:"file"`
	Mode string `json:"mode"`
}

func Flip(c *gin.Context) {
	var flipST FlipST
	if err := c.ShouldBindJSON(&flipST); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON request"})
		return
	}
	file, err := os.Open(flipST.File)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File not found"})
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"response": "Failed to properly close the file"})
		}
	}(file)
	srcImage, err := imaging.Decode(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image format"})
		return
	}
	var dstImage image.Image
	switch flipST.Mode {
	case "horizontal":
		dstImage = imaging.FlipH(srcImage)
	case "vertical":
		dstImage = imaging.FlipV(srcImage)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Mode must be 'horizontal' or 'vertical'"})
		return
	}

	var buf bytes.Buffer
	opts := &jpeg.Options{Quality: 100}
	if err = jpeg.Encode(&buf, dstImage, opts); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode image"})
		return
	}

	c.Data(http.StatusOK, "image/jpeg", buf.Bytes())

}
