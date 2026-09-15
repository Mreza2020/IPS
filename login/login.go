package Login

import "github.com/Mreza2020/Image_Processing_Service/DB"

// Login authenticates a user by verifying the provided username and password
// against the stored user records. It returns "ok" and the authenticated user
// when the credentials are valid, or an empty string and nil when authentication
// fails.
func Login(username string, password string) (string, *DB.User) {
	if username == "" || password == "" {

		return "", nil
	}

	for _, user := range DB.UserStorage {
		if user.Username == username && user.Password == password {

			return "ok", &user
		}
	}

	return "", nil
}
