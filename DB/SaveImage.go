package DB

import (
	"bytes"
	"fmt"
	"os"
)

// SaveImage saves the contents of the provided image buffer to the specified
// output file and returns the output file path. It returns an empty string if
// the image cannot be written to the file.
func SaveImage(buf *bytes.Buffer, outputFile string) string {
	err := os.WriteFile(outputFile, buf.Bytes(), 0644)
	if err != nil {
		fmt.Printf("Failed to save image: %s\n", err)
		return ""
	}

	return outputFile
}
