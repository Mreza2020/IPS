package DB

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	SerializeMode string
	UserStorage   []User
)

const (
	Path           = "text.txt"
	SerializeMode1 = "txt"
	SerializeMode2 = "json"
)

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// WriteUsers persists user data to the configured storage file according to
// the selected serialization mode. SerializeMode1 appends the provided data,
// while SerializeMode2 serializes the current UserStorage collection as
// formatted JSON and overwrites the existing file.
func WriteUsers(data string) {
	switch SerializeMode {

	case SerializeMode1:
		file, err := os.OpenFile(Path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		_, err = file.WriteString(data)
		if err != nil {
			fmt.Println(err)
		}

	case SerializeMode2:
		data, err := json.MarshalIndent(UserStorage, "", "    ")
		if err != nil {
			fmt.Println(err)
			return
		}

		file, err := os.OpenFile(Path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		_, err = file.Write(data)
		if err != nil {
			fmt.Println(err)
		}

	default:
		fmt.Println("unknown serialize mode:", SerializeMode)
	}
}

// LoadUsers reads user records from the configured storage file, clears the
// current in-memory user list, and deserializes the data according to the
// selected serialization mode. It validates text-based records when using
// SerializeMode1 and decodes JSON data when using SerializeMode2.
func LoadUsers() {
	file, err := os.Open(Path)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}

		fmt.Println(err)
		return
	}
	defer file.Close()

	UserStorage = nil

	switch SerializeMode {
	case SerializeMode1:
		scanner := bufio.NewScanner(file)

		lineNumber := 0

		for scanner.Scan() {
			lineNumber++

			line := scanner.Text()

			parts := strings.Split(line, ", ")

			if len(parts) != 3 {
				fmt.Printf(
					"Invalid data at line %d: %s\n",
					lineNumber,
					line,
				)
				continue
			}

			if !strings.HasPrefix(parts[0], "id: ") ||
				!strings.HasPrefix(parts[1], "name: ") ||
				!strings.HasPrefix(parts[2], "password: ") {

				fmt.Printf(
					"Invalid format at line %d: %s\n",
					lineNumber,
					line,
				)
				continue
			}

			id := strings.TrimPrefix(parts[0], "id: ")
			name := strings.TrimPrefix(parts[1], "name: ")
			password := strings.TrimPrefix(parts[2], "password: ")

			idU, err := strconv.ParseUint(id, 10, 64)
			if err != nil {
				fmt.Printf(
					"Invalid ID at line %d: %v\n",
					lineNumber,
					err,
				)
				continue
			}

			user := User{
				ID:       idU,
				Username: name,
				Password: password,
			}

			UserStorage = append(UserStorage, user)
		}

		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}

	case SerializeMode2:
		data, err := io.ReadAll(file)
		if err != nil {
			fmt.Println(err)
			return
		}

		if len(data) == 0 {
			return
		}

		err = json.Unmarshal(data, &UserStorage)
		if err != nil {
			fmt.Println(err)
			return
		}

	default:
		fmt.Println("unknown serialize mode:", SerializeMode)
	}
}
