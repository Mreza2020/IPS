package Build

import (
	"bytes"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"image"
	"image/color"
	"image/jpeg"
	"net/http"
	"os"
	"strings"
)

type FilterRequest struct {
	File   string `json:"file"`
	Filter string `json:"filter"`
}

func ApplyFilter(c *gin.Context) {
	var filterRequest FilterRequest
	if err := c.ShouldBindJSON(&filterRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON request"})
		return
	}
	file, err := os.Open(filterRequest.File)
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
	filter := strings.ToLower(filterRequest.Filter)
	var dstImage image.Image
	switch filter {
	case "grayscale":
		dstImage = imaging.Grayscale(srcImage)
	case "sepia":
		dstImage = imaging.AdjustFunc(srcImage, func(c color.NRGBA) color.NRGBA {
			r := float64(c.R)
			g := float64(c.G)
			b := float64(c.B)
			tr := 0.393*r + 0.769*g + 0.189*b
			tg := 0.349*r + 0.686*g + 0.168*b
			tb := 0.272*r + 0.534*g + 0.131*b

			if tr > 255 {
				tr = 255
			}
			if tg > 255 {
				tg = 255
			}
			if tb > 255 {
				tb = 255
			}
			return color.NRGBA{R: uint8(tr), G: uint8(tg), B: uint8(tb), A: c.A}
		})
	case "invert":
		dstImage = imaging.Invert(srcImage)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported filter"})
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
