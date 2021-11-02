package main

import (
	"crypto/rand"
	"log"

	"golang.org/x/crypto/argon2"
)

type options struct {
	saltLength  uint32
	time        uint32
	memory      uint32
	parallelism uint8
	keyLen      uint32
}

// Generate an Argon2 hash with specifc parameters defined in the options
func generateArgon2Hash(data []byte) []byte {
	config := &options{
		saltLength:  16,
		time:        4,
		memory:      64 * 1024,
		parallelism: 2,
		keyLen:      32,
	}

	hash, err := hasher(data, config)
	if err != nil {
		log.Fatal(err)
	}

	return hash
}

func hasher(data []byte, config *options) (hash []byte, err error) {
	// generate a cryptographically secure random salt
	salt, err := salter(config.saltLength)
	if err != nil {
		return nil, err
	}

	// using the argon2 ID variant, hash the data w/ the randomized salt parameters
	hash = argon2.IDKey(data,
		salt,
		config.time,
		config.memory,
		config.parallelism,
		config.keyLen,
	)

	return hash, nil
}

func salter(n uint32) ([]byte, error) {
	arr := make([]byte, n)
	_, err := rand.Read(arr)
	if err != nil {
		return nil, err
	}

	return arr, nil
}
