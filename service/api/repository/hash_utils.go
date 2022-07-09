package repository

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
)

// In order to store passwords securely in DB, passwords will be salted and hashed

const SALT_SIZE = 16 // Bytes size

func GenerateRandomSalt() ([]byte, error) {
	salt := make([]byte, SALT_SIZE)

	_, err := rand.Read(salt)
	if err != nil {
		return nil, fmt.Errorf("failed generating salt")
	}

	return salt, nil
}

func HashPassword(password string, salt []byte) string {
	passwordBytes := []byte(password)
	passwordBytes = append(passwordBytes, salt...)

	sha512Hasher := sha512.New()
	sha512Hasher.Write(passwordBytes)
	hashedPasswordBytes := sha512Hasher.Sum(nil)

	base64EncodedHasePassword := base64.URLEncoding.EncodeToString(hashedPasswordBytes)

	return base64EncodedHasePassword

}

func DoPasswordsMatch(password string, salt []byte, hashedPassword string) bool {
	currHashed := HashPassword(password, salt)
	return currHashed == hashedPassword
}
