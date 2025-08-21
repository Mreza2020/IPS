package Build

import (
	"bytes"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"image/color"
	"image/jpeg"
	"net/http"
	"os"
	"strconv"
)

type RotateST struct {
	Rotate string `json:"rotate"`
	File   string `json:"file"`
}

func Rotate(c *gin.Context) {
	var rotate RotateST
	if err := c.ShouldBindJSON(&rotate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON request"})
	}
	angle, err := strconv.ParseFloat(rotate.Rotate, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid angle"})
		return
	}
	srcFile, err := os.Open(rotate.File)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File not found"})
	}
	defer func(srcFile *os.File) {
		err = srcFile.Close()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"response": "Failed to properly close the file"})
			return
		}
	}(srcFile)

	srcImage, err := imaging.Decode(srcFile)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image format"})
		return
	}
	rotated := imaging.Rotate(srcImage, angle, color.Black)

	var buf bytes.Buffer
	opts := &jpeg.Options{Quality: 100}

	if err = jpeg.Encode(&buf, rotated, opts); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode image"})
		return
	}

	c.Data(http.StatusOK, "image/jpeg", buf.Bytes())
}
