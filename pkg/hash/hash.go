package hash

import (
	"golang.org/x/crypto/bcrypt"
)

// GenerateHash will generate hashed string
func GenerateHash(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// CompareHash will compare pwd (plain) with its hashed string
func CompareHash(pwd string, hashed string) bool {

	byteHash := []byte(hashed)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(pwd))

	return err == nil
}
