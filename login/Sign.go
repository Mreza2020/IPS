package Login

import (
	"fmt"
)

type User struct {
	ID       uint64
	Username string
	Password string
}

var UserStorage []User

var Authentication *User

func Sign(name string, password string) string {
	if len(name) == 0 || len(password) == 0 {

		return ""
	}
	for _, username := range UserStorage {
		if username.Username == name {
			fmt.Println("Username already exists.")

			return ""
		}
	}

	sign := User{
		ID:       uint64(len(UserStorage) + 1),
		Username: name,
		Password: password,
	}

	UserStorage = append(UserStorage, sign)

	return "ok"
}
