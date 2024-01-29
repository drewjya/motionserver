package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

// Hash the given bytes using SHA256.
func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// ValidateHash validate whether the given password is match with the given hash.
func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
