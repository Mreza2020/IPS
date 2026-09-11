package Build

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/Mreza2020/Image_Processing_Service/DB"
	"github.com/disintegration/imaging"
)

// Resize scales the specified image to the given width and height using the
// Lanczos resampling algorithm and saves the resized image as a new file.
// Supported image formats are JPEG and PNG. It returns "ok" when the operation
// succeeds and an empty string if the image is invalid, the format is
// unsupported, or the resized image cannot be encoded or saved.
func Resize(fileName, width, height string) string {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer file.Close()

	srcImage, err := imaging.Decode(file)
	if err != nil {
		fmt.Printf("Invalid image format : %v", err)

		return ""
	}
	w := ConvertToInt(width)
	h := ConvertToInt(height)

	dstImage := imaging.Resize(srcImage, w, h, imaging.Lanczos)

	var buf bytes.Buffer
	var outputFile string

	ext := strings.ToLower(filepath.Ext(fileName))

	fileName1 := strings.TrimSuffix(filepath.Base(fileName), filepath.Ext(fileName))

	dir := filepath.Dir(fileName)
	switch ext {
	case ".jpg", ".jpeg":
		opts := &jpeg.Options{Quality: 100}
		if err = jpeg.Encode(&buf, dstImage, opts); err != nil {
			fmt.Printf("Failed to encode image : %v", err)

			return ""
		}
		outputFile = filepath.Join(dir, fileName1+"_resize.jpg")
	case ".png":
		if err = png.Encode(&buf, dstImage); err != nil {
			fmt.Printf("Failed to encode image: %s\n", err)

			return ""

		}
		outputFile = filepath.Join(dir, fileName1+"_resize.png")
	default:
		fmt.Println("Unsupported file extension")

		return ""

	}
	path := DB.SaveImage(&buf, outputFile)
	fmt.Printf("Resize %s to %s\n", fileName, path)

	return "ok"

}
