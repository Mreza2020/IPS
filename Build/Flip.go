package Build

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/Mreza2020/Image_Processing_Service/DB"
	"github.com/disintegration/imaging"
)

// Flip flips the specified image horizontally or vertically and saves the
// transformed image as a new file while preserving the original image format.
// It returns "ok" when the operation succeeds and an empty string if the image
// is invalid, the flip mode or file format is unsupported, or encoding or
// saving the processed image fails.
func Flip(fileName, mode string) string {
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
	var dstImage image.Image
	switch strings.ToLower(mode) {
	case "horizontal":
		dstImage = imaging.FlipH(srcImage)
	case "vertical":
		dstImage = imaging.FlipV(srcImage)
	default:
		fmt.Println("Mode must be 'horizontal' or 'vertical'")

		return ""
	}

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
		outputFile = filepath.Join(dir, fileName1+"_Flip.jpg")
	case ".png":
		if err = png.Encode(&buf, dstImage); err != nil {
			fmt.Printf("Failed to encode image: %s\n", err)

			return ""

		}
		outputFile = filepath.Join(dir, fileName1+"_Flip.png")
	default:
		fmt.Println("Unsupported file extension")

		return ""

	}
	path := DB.SaveImage(&buf, outputFile)
	fmt.Printf("Flip %s applied to %s\n", mode, path)

	return "ok"
}
