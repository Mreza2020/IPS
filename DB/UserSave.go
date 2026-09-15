package DB

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const Path = "text.txt"

var File *os.File

type User struct {
	ID       uint64
	Username string
	Password string
}

var UserStorage []User

// init initializes the user storage file and loads the existing users into memory.
func init() {
	file, err := os.OpenFile(Path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
	}

	File = file

	LoadUsers()
}

// WriteUsers appends the provided user data to the persistent storage file.
func WriteUsers(data string) {
	_, err := File.WriteString(data)
	if err != nil {
		fmt.Println(err)

		return
	}
}

// LoadUsers reads user records from the storage file and loads them into
// the in-memory UserStorage collection.
func LoadUsers() {
	scanner := bufio.NewScanner(File)

	lineNumber := 0

	for scanner.Scan() {
		lineNumber++

		line := scanner.Text()

		parts := strings.Split(line, ", ")

		if len(parts) != 3 {
			fmt.Printf("Invalid data at line %d: %s\n", lineNumber, line)

			continue
		}

		if !strings.HasPrefix(parts[0], "id: ") || !strings.HasPrefix(parts[1], "name: ") || !strings.HasPrefix(parts[2], "password: ") {
			fmt.Printf("Invalid format at line %d: %s\n", lineNumber, line)

			continue
		}

		id := strings.TrimPrefix(parts[0], "id: ")
		name := strings.TrimPrefix(parts[1], "name: ")
		password := strings.TrimPrefix(parts[2], "password: ")

		idU, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			fmt.Println(err)
		}

		sign := User{
			ID:       idU,
			Username: name,
			Password: password,
		}

		UserStorage = append(UserStorage, sign)

	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
