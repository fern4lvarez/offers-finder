package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestNewToken(t *testing.T) {
	token := NewToken()
	if len(token.Key) != 40 {
		t.Errorf(errorMessage, "NewToken", 40, len(token.Key))
	}
}

func TestTokenDecode(t *testing.T) {
	b := []byte(`{"token":"abcdefghijklmnopqrstuvxyz"}`)
	expectedToken := Token{
		Key: "abcdefghijklmnopqrstuvxyz",
	}

	var token Token
	err := token.Decode(b)

	if err != nil {
		t.Errorf(errorMessage, "TokenDecode", nil, err)
	}
	if token.Key != "abcdefghijklmnopqrstuvxyz" {
		t.Errorf(errorMessage, "TokenDecode", token, expectedToken.Key)
	}
}

func TestTokenEncode(t *testing.T) {
	token := Token{
		Key: "abcdefghijklmnopqrstuvxyz",
	}
	expectedEncoded := []byte(`{"token":"abcdefghijklmnopqrstuvxyz"}`)

	b, err := token.Encode()

	if err != nil {
		t.Errorf(errorMessage, "TokenEncode", nil, err)
	}
	if bytes.Compare(b, expectedEncoded) != 0 {
		t.Errorf(errorMessage, "TokenEncode", string(b), string(expectedEncoded))
	}
}

func ExampleTokenString() {
	token := Token{
		Key: "abcdefghijklmnopqrstuvxyz",
	}
	fmt.Println(token)
	// Output: {"token":"abcdefghijklmnopqrstuvxyz"}
}

func TestGenerateToken(t *testing.T) {
	token := generateToken(-1)
	if len(token) != 10 {
		t.Errorf(errorMessage, "generateToken", 28, len(token))
	}

	token = generateToken(28)
	if len(token) != 28 {
		t.Errorf(errorMessage, "genereateToken", 28, len(token))
	}
}
