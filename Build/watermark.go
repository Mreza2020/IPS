package Build

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/Mreza2020/Image_Processing_Service/DB"
	"github.com/disintegration/imaging"
)

// Watermark applies a resized watermark to the specified image with the given
// opacity and display mode, then saves the processed image as a new file.
// The watermark can be placed once in the bottom-right corner or repeated
// across the image in a tiled layout. Supported output formats are JPEG and PNG.
// It returns "ok" when the operation succeeds and an empty string if the input
// images are invalid, the opacity or mode is invalid, the format is unsupported,
// or the processed image cannot be encoded or saved.
func Watermark(fileName, fileNameW, opacity, mode string) string {
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

	fileW, err := os.Open(fileNameW)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer fileW.Close()

	srcImage1, err := imaging.Decode(fileW)
	if err != nil {
		fmt.Printf("Invalid image format : %v", err)

		return ""
	}

	val := ConvertToInt(opacity)
	opacityD := 50
	if val >= 0 && val <= 100 {
		opacityD = val
	}

	mode1 := strings.ToLower(mode)
	if mode1 != "single" && mode1 != "tile" {
		fmt.Println("Mode must be 'single' or 'tile'")
		return ""
	}

	Resized := imaging.Resize(srcImage1, srcImage.Bounds().Dx()/5, 0, imaging.Lanczos)

	WithOpacity := imaging.AdjustFunc(Resized, func(c color.NRGBA) color.NRGBA {
		newA := uint8(float64(c.A) * float64(opacityD) / 100)
		return color.NRGBA{R: c.R, G: c.G, B: c.B, A: newA}
	})

	dst := imaging.Clone(srcImage)

	if mode1 == "single" {
		offset := image.Pt(dst.Bounds().Dx()-WithOpacity.Bounds().Dx()-10, dst.Bounds().Dy()-WithOpacity.Bounds().Dy()-10)

		draw.Draw(dst, WithOpacity.Bounds().Add(offset), WithOpacity, image.Point{}, draw.Over)
	}
	if mode1 == "tile" {
		watermarkWidth := WithOpacity.Bounds().Dx()
		watermarkHeight := WithOpacity.Bounds().Dy()
		if watermarkWidth <= 0 || watermarkHeight <= 0 {
			fmt.Println("Invalid watermark dimensions")
			return ""
		}

		gapX := 30
		gapY := 30

		for y := 0; y < dst.Bounds().Dy(); y += watermarkHeight + gapY {
			for x := 0; x < dst.Bounds().Dx(); x += watermarkWidth + gapX {
				offset := image.Pt(x, y)
				draw.Draw(dst, WithOpacity.Bounds().Add(offset), WithOpacity, image.Point{}, draw.Over)
			}
		}
	}

	var buf bytes.Buffer
	var outputFile string

	ext := strings.ToLower(filepath.Ext(fileName))

	fileName1 := strings.TrimSuffix(filepath.Base(fileName), filepath.Ext(fileName))

	dir := filepath.Dir(fileName)
	switch ext {
	case ".jpg", ".jpeg":
		opts := &jpeg.Options{Quality: 100}
		if err = jpeg.Encode(&buf, dst, opts); err != nil {
			fmt.Printf("Failed to encode image : %v", err)

			return ""
		}
		outputFile = filepath.Join(dir, fileName1+"_Watermark.jpg")
	case ".png":
		if err = png.Encode(&buf, dst); err != nil {
			fmt.Printf("Failed to encode image: %s\n", err)

			return ""

		}
		outputFile = filepath.Join(dir, fileName1+"_Watermark.png")
	default:
		fmt.Println("Unsupported file extension")

		return ""

	}
	path := DB.SaveImage(&buf, outputFile)
	fmt.Printf("Watermark applied to %s\n", path)

	return "ok"

}
