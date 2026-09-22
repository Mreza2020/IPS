package DB

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

type Image struct {
	User string
	Path string
}

var ImageDB []Image

const path = "txt-images.txt"

var outputpath string

// SaveImage saves the contents of the provided image buffer to the specified
// output file and returns the output file path. It returns an empty string if
// the image cannot be written to the file.
func SaveImage(buf *bytes.Buffer, outputFile string) string {
	err := os.WriteFile(outputFile, buf.Bytes(), 0644)
	if err != nil {
		fmt.Printf("Failed to save image: %s\n", err)

		return ""
	}
	outputpath = outputFile

	return outputFile
}

// SaveImageDB records the user and output file path of a processed image in
// the image history file. It appends the record to the existing file and
// returns without writing if the file cannot be opened or updated.
func SaveImageDB(user string) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)

		return
	}
	defer file.Close()

	data := fmt.Sprintf("Created by %v, File path %v\n", user, outputpath)
	_, err = file.WriteString(data)
	if err != nil {
		fmt.Println(err)

		return
	}
}

// LoadImageDB reads image records from the image history file, validates each
// record's format, and appends valid entries to the in-memory ImageDB collection.
// Invalid records are skipped and reported with their corresponding line number.
func LoadImageDB() {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	lineNumber := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()

		parts := strings.Split(line, ", ")

		if len(parts) != 2 {
			fmt.Printf(
				"Invalid data at line %d: %s\n",
				lineNumber,
				line,
			)
			continue
		}

		if !strings.HasPrefix(parts[0], "Created by ") || !strings.Contains(parts[1], "File path ") {
			fmt.Printf("Invalid format at line %d: %s\n", lineNumber, line)

			continue
		}

		user := strings.TrimPrefix(parts[0], "Created by ")
		pathF := strings.TrimPrefix(parts[1], "File path ")

		image := Image{
			User: user,
			Path: pathF,
		}

		ImageDB = append(ImageDB, image)
	}
}
