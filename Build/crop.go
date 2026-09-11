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

// Crop extracts a rectangular region from the specified image using the given
// width, height, and starting X/Y coordinates, then saves the cropped image as
// a new file while preserving the original image format. It returns "ok" when
// the operation succeeds and an empty string if the input image is invalid,
// the crop region is outside the image bounds, encoding fails, or the format
// is unsupported.
func Crop(width, height, x, y, fileName string) string {
	w := ConvertToInt(width)
	h := ConvertToInt(height)
	x1 := ConvertToInt(x)
	y1 := ConvertToInt(y)

	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer f.Close()

	srcImage, err := imaging.Decode(f)
	if err != nil {
		fmt.Printf("Invalid image format : %v", err)

		return ""
	}
	if x1 < 0 || y1 < 0 ||
		w <= 0 || h <= 0 ||
		x1+w > srcImage.Bounds().Dx() ||
		y1+h > srcImage.Bounds().Dy() {
		fmt.Println("Crop dimensions out of bounds")
		return ""
	}
	cropped := imaging.Crop(srcImage, image.Rect(x1, y1, x1+w, y1+h))

	var buf bytes.Buffer
	var outputFile string

	ext := strings.ToLower(filepath.Ext(fileName))

	fileName1 := strings.TrimSuffix(filepath.Base(fileName), filepath.Ext(fileName))

	dir := filepath.Dir(fileName)
	switch ext {
	case ".jpg", ".jpeg":
		opts := &jpeg.Options{Quality: 100}
		if err = jpeg.Encode(&buf, cropped, opts); err != nil {
			fmt.Printf("Failed to encode image : %v", err)

			return ""
		}
		outputFile = filepath.Join(dir, fileName1+"_cropped.jpg")
	case ".png":
		if err = png.Encode(&buf, cropped); err != nil {
			fmt.Printf("Failed to encode image: %s\n", err)

			return ""

		}
		outputFile = filepath.Join(dir, fileName1+"_cropped.png")
	default:
		fmt.Println("Unsupported file extension")

		return ""

	}
	path := DB.SaveImage(&buf, outputFile)
	fmt.Printf("Crop %s to %s\n", fileName, path)

	return "ok"

}
