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

// ChangeFormat converts the specified image to the requested format and saves
// the converted image as a new file. Supported output formats are JPEG and PNG.
// The function returns "ok" when the conversion succeeds and an empty string
// if the input image is invalid, the requested format is unsupported, or the
// converted image cannot be encoded or saved.
func ChangeFormat(fileName, format string) string {
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
	var buf bytes.Buffer
	var outputFile string

	format1 := strings.ToLower(strings.TrimPrefix(format, "."))

	fileName1 := strings.TrimSuffix(filepath.Base(fileName), filepath.Ext(fileName))

	dir := filepath.Dir(fileName)
	switch format1 {
	case "jpg", "jpeg":
		opts := &jpeg.Options{Quality: 100}
		if err = jpeg.Encode(&buf, srcImage, opts); err != nil {
			fmt.Printf("Failed to encode image : %v", err)

			return ""
		}
		outputFile = filepath.Join(dir, fileName1+"Format.jpg")
	case "png":
		if err = png.Encode(&buf, srcImage); err != nil {
			fmt.Printf("Failed to encode image: %s\n", err)

			return ""

		}
		outputFile = filepath.Join(dir, fileName1+"_Format.png")
	default:
		fmt.Println("Unsupported file extension")

		return ""

	}
	path := DB.SaveImage(&buf, outputFile)
	fmt.Printf("Format changed to %s: %s -> %s\n", format1, fileName, path)

	return "ok"
}
