package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"net/http"

	"github.com/satori/go.uuid"
	"golang.org/x/crypto/argon2"
)

func HandleRegister(input *RegisterInput) *ErrorOutput {
	// Check if the duplicated user information exists.
	var user User
	DB.Where(&User{Username: input.Username}).First(&user)
	if user.UUID != "" {
		return &ErrorOutput{http.StatusBadRequest, "Username already exists."}
	}
	DB.Where(&User{Email: input.Email}).First(&user)
	if user.UUID != "" {
		return &ErrorOutput{http.StatusBadRequest, "Email already exists."}
	}

	// Create new user.
	salt := make([]byte, 32)
	rand.Read(salt)
	user = User{
		UUID:         uuid.NewV4().String(),
		Username:     input.Username,
		Email:        input.Email,
		PasswordHash: argon2.IDKey([]byte(input.Password), salt, 1, 8*1024, 4, 32),
		PasswordSalt: salt,
	}

	errs := DB.Save(&user).GetErrors()
	if len(errs) > 0 {
		fmt.Errorf("%v", errs)
		return &ErrorOutput{http.StatusInternalServerError, "Server error :(\nPlease try again."}
	}

	// TODO - Send activation email.

	return nil
}

func HandleTokenGenerate(input *TokenGenerateInput) (*TokenGenerateOutput, *ErrorOutput) {
	// Check login info.
	var user User
	DB.Where(&User{Email: input.Email}).First(&user)
	if user.UUID == "" {
		return nil, &ErrorOutput{http.StatusUnauthorized, "Invalid email or password."}
	}
	hash := argon2.IDKey([]byte(input.Password), user.PasswordSalt, 1, 8*1024, 4, 32)
	if bytes.Compare(hash, user.PasswordHash) != 0 {
		return nil, &ErrorOutput{http.StatusUnauthorized, "Invalid email or password."}
	}

	// Create new access token.
	token, err := NewToken(user.UUID)
	if err != nil {
		return nil, &ErrorOutput{http.StatusInternalServerError, "Sign in failed."}
	}

	DB.Save(&token)
	return &TokenGenerateOutput{AccessToken: token.AccessToken, RefreshToken: token.RefreshToken}, nil
}
