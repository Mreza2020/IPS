package Security

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword securely hashes the provided password using bcrypt and returns
// the resulting hash. It returns an empty string if password hashing fails.
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(bytes)
}

// CheckPasswordHash verifies the provided password against a bcrypt hash and
// returns true if the password matches or false if the verification fails.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
