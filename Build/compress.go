package Build

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Mreza2020/Image_Processing_Service/DB"
	"github.com/disintegration/imaging"
)

// Compress decodes the specified image, re-encodes it according to its format,
// and saves the compressed result as a new file. For JPEG images, the quality
// is controlled by the provided value from 1 to 100. PNG images are re-encoded
// using the standard PNG encoder. The function returns "ok" on success and an
// empty string if the operation fails or the image format is unsupported.
func Compress(file, quality string) string {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer f.Close()

	srcImage, err := imaging.Decode(f)
	if err != nil {
		fmt.Printf("Invalid image format: %s\n", err)
		return ""
	}

	DefaultQuality := 75
	if val, err1 := strconv.Atoi(quality); err1 == nil && val >= 1 && val <= 100 {
		DefaultQuality = val
	} else {
		fmt.Println("Invalid quality")
	}

	ext := strings.ToLower(filepath.Ext(file))

	fileName := strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))

	dir := filepath.Dir(file)

	var buf bytes.Buffer
	var outputFile string

	switch ext {
	case ".jpg", ".jpeg":
		if err = jpeg.Encode(&buf, srcImage, &jpeg.Options{Quality: DefaultQuality}); err != nil {
			fmt.Printf("Failed to encode image: %s\n", err)

			return ""

		}
		outputFile = filepath.Join(dir, fileName+"_compressd.jpg")
	case ".png":
		if err = png.Encode(&buf, srcImage); err != nil {
			fmt.Printf("Failed to encode image: %s\n", err)

			return ""

		}
		outputFile = filepath.Join(dir, fileName+"_compressd.png")
	default:
		fmt.Println("Unsupported file extension")

		return ""

	}
	path := DB.SaveImage(&buf, outputFile)
	fmt.Printf("Compressed %s to %s\n", file, path)

	return "ok"

}
