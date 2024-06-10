package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
)

func generateAPIKey() (string, error) {
	// Create a byte slice with a length of 32 bytes (256 bits)
	key := make([]byte, 32)

	// Fill the byte slice with random bytes from the cryptographic random number generator
	_, err := rand.Read(key)
	if err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %v", err)
	}

	// Encode the byte slice to a base64 string
	apiKey := base64.StdEncoding.EncodeToString(key)

	return apiKey, nil
}

func main() {
	apiKey, err := generateAPIKey()
	if err != nil {
		log.Fatalf("Error generating API key: %v", err)
	}

	fmt.Println("Generated API Key:", apiKey)
}
