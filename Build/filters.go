package Build

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/Mreza2020/Image_Processing_Service/DB"
	"github.com/disintegration/imaging"
)

// ApplyFilter applies the specified filter to an image and saves the processed
// result as a new file while preserving the original image format. Supported
// filters are grayscale, sepia, and invert. It returns "ok" when the operation
// succeeds and an empty string if the image is invalid, the filter or file
// format is unsupported, or the processed image cannot be encoded or saved.
func ApplyFilter(fileName, filter string) string {
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

	filter1 := strings.ToLower(filter)
	var dstImage image.Image
	switch filter1 {
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
		fmt.Printf("Unsupported filter: %s\n", filter)

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
		outputFile = filepath.Join(dir, fileName1+"_filter.jpg")
	case ".png":
		if err = png.Encode(&buf, dstImage); err != nil {
			fmt.Printf("Failed to encode image: %s\n", err)

			return ""

		}
		outputFile = filepath.Join(dir, fileName1+"_filter.png")
	default:
		fmt.Println("Unsupported file extension")

		return ""

	}
	path := DB.SaveImage(&buf, outputFile)
	fmt.Printf("Filter %s applied to %s\n", filter1, path)

	return "ok"
}
