package Build

import (
	"bytes"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"image"
	"image/jpeg"
	"net/http"
	"os"
	"strconv"
)

type CropS struct {
	Width  string `json:"width"`
	Height string `json:"height"`
	X      string `json:"x"`
	Y      string `json:"y"`
	File   string `json:"file"`
}

func Crop(c *gin.Context) {
	var crop CropS
	if err := c.ShouldBindJSON(&crop); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON request"})
		return
	}
	x, err := strconv.Atoi(crop.X)
	y, err := strconv.Atoi(crop.Y)
	if err != nil || x < 0 || y < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"response": "!!!The file is not valid!!!(x or y)"})
		return
	}
	width, err := strconv.Atoi(crop.Width)
	height, err := strconv.Atoi(crop.Height)
	if err != nil || width < 0 || height < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"response": "!!!The file is not valid!!!(Width or Height)"})
		return
	}
	file, err := os.Open(crop.File)
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
	if x+width > srcImage.Bounds().Dx() || y+height > srcImage.Bounds().Dy() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Crop dimensions out of bounds"})
		return
	}
	cropped := imaging.Crop(srcImage, image.Rect(x, y, x+width, y+height))
	var buf bytes.Buffer
	opts := &jpeg.Options{Quality: 100}
	if err = jpeg.Encode(&buf, cropped, opts); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode image"})
		return
	}

	c.Data(http.StatusOK, "image/jpeg", buf.Bytes())

}
