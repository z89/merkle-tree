package main

import (
	"crypto/rand"
	"log"

	"golang.org/x/crypto/argon2"
)

type options struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

// Generate an Argon2 hash with specifc parameters defined in the options
func generateHash(data string) []byte {
	config := &options{
		memory:      64 * 1024,
		iterations:  5,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}

	hash, err := hasher(data, config)
	if err != nil {
		log.Fatal(err)
	}

	return hash
}

func hasher(data string, config *options) (hash []byte, err error) {
	// generate a cryptographically secure random salt
	salt, err := salter(config.saltLength)
	if err != nil {
		return nil, err
	}

	// using the argon2 ID variant, hash the data w/ the randomized salt parameters
	hash = argon2.IDKey([]byte(data),
		salt,
		config.iterations,
		config.memory,
		config.parallelism,
		config.keyLength,
	)

	return hash, nil
}

func salter(saltLength uint32) ([]byte, error) {
	salt := make([]byte, saltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	} else {
		return salt, nil
	}
}
