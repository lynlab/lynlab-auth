package main

import (
	"fmt"
	"testing"
)

func TestNewToken(t *testing.T) {
	userUUID := "00000000-0000-0000-0000-000000000000"
	token, err := NewToken(userUUID)

	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if token.AccessToken == "" || token.RefreshToken == "" {
		fmt.Println("Generated tokens are invalid")
		t.Fail()
	}
}

func TestValidate(t *testing.T) {
	userUUID := "00000000-0000-0000-0000-000000000000"
	token, _ := NewToken(userUUID)

	parsedUUID, err := token.Validate()

	if err != nil || parsedUUID != userUUID {
		fmt.Println("Token validation failed")
		t.Fail()
	}
}
