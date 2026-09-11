package Build

import (
	"bytes"
	"fmt"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/Mreza2020/Image_Processing_Service/DB"
	"github.com/disintegration/imaging"
)

// Rotate rotates the specified image by the given angle in degrees using a
// black background for empty areas, then saves the rotated image as a new file.
// Supported image formats are JPEG and PNG. It returns "ok" when the operation
// succeeds and an empty string if the image is invalid, the format is unsupported,
// or the rotated image cannot be encoded or saved.
func Rotate(fileName, rotate string) string {
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

	rotated := imaging.Rotate(srcImage, ConvertToFloat(rotate), color.Black)

	var buf bytes.Buffer
	var outputFile string

	ext := strings.ToLower(filepath.Ext(fileName))

	fileName1 := strings.TrimSuffix(filepath.Base(fileName), filepath.Ext(fileName))

	dir := filepath.Dir(fileName)
	switch ext {
	case ".jpg", ".jpeg":
		opts := &jpeg.Options{Quality: 100}
		if err = jpeg.Encode(&buf, rotated, opts); err != nil {
			fmt.Printf("Failed to encode image : %v", err)

			return ""
		}
		outputFile = filepath.Join(dir, fileName1+"_Rotate.jpg")
	case ".png":
		if err = png.Encode(&buf, rotated); err != nil {
			fmt.Printf("Failed to encode image: %s\n", err)

			return ""

		}
		outputFile = filepath.Join(dir, fileName1+"_Rotate.png")
	default:
		fmt.Println("Unsupported file extension")

		return ""

	}
	path := DB.SaveImage(&buf, outputFile)
	fmt.Printf("Rotate %s° applied to %s\n", rotate, path)

	return "ok"

}
