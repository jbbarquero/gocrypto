package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

// Notes:
// See https://www.packtpub.com/networking-and-servers/mastering-go chapter 6
// See https://golang.org/pkg/crypto/cipher/#NewGCM (expand examples)
// See Java Cryptography: Tools and Techniques chapters 2 and 4

func toByteArray(s string) []byte {
	bytes := make([]byte, len(s))

	for i := 0; i < len(bytes); i++ {
		bytes[i] = byte(s[i])
	}

	return bytes
}

func EncryptWithGCMBlockMode(data []byte, key []byte) ([]byte, error) {

	//Initialize the block cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//Generate a randomized nonce
	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	nonce, _ = hex.DecodeString("bbaa99887766554433221100") //For comparing with the Java code

	//Choose a block cipher mode
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return aesgcm.Seal(nil, nonce, data, nil), nil

}

func DecryptWithGCMBlockMode(ciphertext []byte, key []byte) ([]byte, error) {
	//Initialize the block cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//The same nonce as before (we'll consider a parameter instead)
	nonce, _ := hex.DecodeString("bbaa99887766554433221100") //For comparing with the Java code

	//Choose a block cipher mode
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return aesgcm.Open(nil, nonce, ciphertext, nil)

}

func main() {
	key, _ := hex.DecodeString("000102030405060708090a0b0c0d0e0f")
	plaintext := toByteArray("hello, world!")
	fmt.Printf("plain: %s\n", hex.EncodeToString(plaintext))

	ciphertext, err := EncryptWithGCMBlockMode(plaintext, key)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("encrypted: %x\n", ciphertext)

	decryptedtext, err := DecryptWithGCMBlockMode(ciphertext, key)

	fmt.Printf("decrypted: %x\n", decryptedtext)
	fmt.Printf("%s\n", string(decryptedtext))

}
