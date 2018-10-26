package hash

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"log"
)

//TODO: do better
var secretKey = "neictr98y85klfgneghre"

// Create a salt string with 32 bytes of crypto/rand data
func generateSalt() string {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatal(err)
	}
	return base64.URLEncoding.EncodeToString(randomBytes)
}

// Hash a password with the salt
func hashPassword(password string, salt string) string {
	hash := hmac.New(sha256.New, []byte(secretKey))
	io.WriteString(hash, password+salt)
	hashedValue := hash.Sum(nil)
	return hex.EncodeToString(hashedValue)
}

func HashPassword(password string) {
	fmt.Println("HASH PASSWORD: ", password)

	salt := generateSalt()

	hashedPassword := hashPassword(password, salt)

	fmt.Println("Password: " + password)
	fmt.Println("Salt: " + salt)
	fmt.Println("Hashed password: " + hashedPassword)
}
