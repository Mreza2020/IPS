package Build

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"strings"
)

type formatRequest struct {
	Format string `json:"format"`
	File   string `json:"file"`
}

func ChangeFormat(c *gin.Context) {
	var req formatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON request"})
		return
	}

	file, err := os.Open(req.File)
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
	srcImage, _, err := image.Decode(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image format"})
		return
	}

	var buf bytes.Buffer
	format := strings.ToLower(req.Format)
	switch format {
	case "jpeg", "jpg":
		if err = jpeg.Encode(&buf, srcImage, &jpeg.Options{Quality: 100}); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode JPEG"})
			return
		}
		c.Data(http.StatusOK, "image/jpeg", buf.Bytes())
	case "png":
		if err = png.Encode(&buf, srcImage); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode PNG"})
			return
		}
		c.Data(http.StatusOK, "image/png", buf.Bytes())
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported format. Use 'jpeg' or 'png'"})
	}

}
