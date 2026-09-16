package Login

import (
	"fmt"

	"github.com/Mreza2020/Image_Processing_Service/DB"
	"github.com/Mreza2020/Image_Processing_Service/Security"
)

var Authentication *DB.User

// Sign registers a new user with the specified username and password, hashes
// the password, stores the user in memory, and persists the account data.
// It returns "ok" when registration succeeds and an empty string if the
// username or password is empty or the username is already registered.
func Sign(name string, password string) string {
	if len(name) == 0 || len(password) == 0 {

		return ""
	}

	for _, username := range DB.UserStorage {
		if username.Username == name {
			fmt.Println("Username already exists.")

			return ""
		}
	}

	sign := DB.User{
		ID:       uint64(len(DB.UserStorage) + 1),
		Username: name,
		Password: Security.HashPassword(password),
	}

	DB.UserStorage = append(DB.UserStorage, sign)

	data := fmt.Sprintf("id: %d, name: %s, password: %s\n", sign.ID, sign.Username, sign.Password)

	DB.WriteUsers(data)

	return "ok"
}
