package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/chacha20poly1305"
)

func main() {
	// Generate a random key and nonce
	key := make([]byte, chacha20poly1305.KeySize)
	nonce := make([]byte, chacha20poly1305.NonceSizeX)
	if _, err := rand.Read(key); err != nil {
		panic(err)
	}
	if _, err := rand.Read(nonce); err != nil {
		panic(err)
	}

	// Plaintext message
	message := []byte("This is a secret message!")

	// Encryption
	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		panic(err)
	}
	ciphertext, err := aead.Seal(nil, nonce, message, nil)
	if err != nil {
		panic(err)
	}

	// Print ciphertext in hexadecimal format
	fmt.Printf("Ciphertext: %s\n", hex.EncodeToString(ciphertext))

	// Decryption (assuming you share the key with the recipient)
	plaintext, err := aead.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println("Error decrypting:", err)
		return
	}

	fmt.Println("Decrypted message:", string(plaintext))
}
