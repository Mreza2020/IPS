package Build

import (
	"bytes"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"net/http"
	"os"
	"strconv"
)

type watermarkST struct {
	Watermark string `json:"watermark"`
	FIle      string `json:"file"`
	Opacity   string `json:"opacity"`
}

func Watermark(c *gin.Context) {
	var watermark watermarkST
	if err := c.ShouldBindJSON(&watermark); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON request"})
		return
	}
	file, err := os.Open(watermark.FIle)
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

	fileW, err := os.Open(watermark.Watermark)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File not found"})
		return
	}
	defer func(fileW *os.File) {
		err = fileW.Close()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"response": "Failed to properly close the file"})
			return
		}
	}(fileW)

	wImage, err := imaging.Decode(fileW)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image format"})
		return
	}
	val, err := strconv.Atoi(watermark.Opacity)
	opacity := 50
	if err == nil && val >= 0 && val <= 100 {
		opacity = val
	}
	Resized := imaging.Resize(wImage, srcImage.Bounds().Dx()/5, 0, imaging.Lanczos)

	WithOpacity := imaging.AdjustFunc(Resized, func(c color.NRGBA) color.NRGBA {
		newA := uint8(float64(c.A) * float64(opacity) / 100)
		return color.NRGBA{R: c.R, G: c.G, B: c.B, A: newA}
	})
	dst := imaging.Clone(srcImage)
	offset := image.Pt(dst.Bounds().Dx()-WithOpacity.Bounds().Dx()-10, dst.Bounds().Dy()-WithOpacity.Bounds().Dy()-10)

	dst1 := imaging.Clone(dst)
	dst2 := imaging.Clone(WithOpacity)

	draw.Draw(dst1, dst2.Bounds().Add(offset), dst2, image.Point{}, draw.Over)

	var buf bytes.Buffer
	opts := &jpeg.Options{Quality: 100}

	if err = jpeg.Encode(&buf, dst1, opts); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode image"})
		return
	}

	c.Data(http.StatusOK, "image/jpeg", buf.Bytes())
}
