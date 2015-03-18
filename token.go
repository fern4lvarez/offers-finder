package main

import (
	"crypto/rand"
	"encoding/json"
	"log"
)

type Token struct {
	Key string `json:"token"`
}

// NewToken returns a pointer to a randomly
// generated Token
func NewToken() *Token {
	return &Token{
		Key: generateToken(40),
	}
}

// Decode decodes json in bytes into a Token
func (token *Token) Decode(b []byte) (err error) {
	return json.Unmarshal(b, &token)
}

// Encode encodes a Token into json in bytes
func (token Token) Encode() (b []byte, err error) {
	b, err = json.Marshal(token)
	return
}

// String formats the output of a Token
func (token Token) String() string {
	b, err := token.Encode()
	if err != nil {
		return ""
	}

	return string(b)
}

// generateToken out of characters with a given size
func generateToken(size int) string {
	if size <= 0 {
		size = 10
	}

	var dictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, size)

	_, err := rand.Read(bytes)
	if err != nil {
		log.Println("WARNING when generating Token.")
	}

	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}

	return string(bytes)
}
