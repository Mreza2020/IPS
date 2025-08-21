package Build

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"image"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func Upload(c *gin.Context) {
	//Login.AuthMiddleware()
	name, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "LastName not found"})
		return
	}
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 6<<20)
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	originalName := file.Filename
	if strings.HasPrefix(originalName, ".") || strings.Contains(originalName, "..") {
		c.String(http.StatusBadRequest, "⚠️ The file name is suspicious.")
		return
	}
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": "!!!Error opening file!!!"})
		return
	}
	defer func(src multipart.File) {
		err = src.Close()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"response": "!!!Error closing file!!!"})
		}
	}(src)

	_, format, err := image.DecodeConfig(src)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": "!!!The file is not valid!!!"})
		return
	}
	allowedFormats := map[string]string{
		"jpeg": ".jpg",
		"png":  ".png",
	}
	ext, ok := allowedFormats[format]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"response": "!!!Format is not allowed!!!"})
		return
	}
	if _, err = src.Seek(0, io.SeekStart); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"response": "!!!Error restoring file!!!"})
		return
	}
	savePath := filepath.Join("C:\\Data", "uploads", "images")
	err = os.MkdirAll(savePath, 0700)
	if err != nil {
		return
	}
	newFileName := uuid.New().String() + ext

	ok1 := DbImage(newFileName, name.(string))
	if ok1 {
		fullPath := filepath.Join(savePath, newFileName)
		dst, err1 := os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY, 0600)
		if err1 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"response": "!!!Error saving image!!!"})

			return
		}
		defer func(dst *os.File) {
			err := dst.Close()
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"response": "!!!Error closing file!!!"})
			}
		}(dst)

		_, err = io.Copy(dst, src)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"response": "!!!Error saving file!!!"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"response": "!!!File uploaded successfully!!!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"response": "!!!File uploaded Error!!!"})
	}

}
