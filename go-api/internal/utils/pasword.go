package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes the given password
func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}

// CheckPassword checks if the given password is correct
func CheckPassword(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
