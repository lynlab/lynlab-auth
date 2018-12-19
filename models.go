package main

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var hmacSecret []byte

/// User
type User struct {
	UUID         string `gorm:"type:varchar(40);pramiry_key"`
	Email        string `gorm:"type:varchar(255);not null;unique_index"`
	Username     string `gorm:"type:varchar(255);not null"`
	PasswordHash []byte `gorm:"not null" json:"-"`
	PasswordSalt []byte `gorm:"not null" json:"-"`
	IsActivated  bool   `gorm:"default:false"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

/// Tokens
type Token struct {
	ID           int
	UserUUID     string `gorm:"type:varchar(40);index"`
	AccessToken  string `gorm:"type:varchar(512);unique_index"`
	RefreshToken string `gorm:"type:varchar(512);unique_index"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type accessTokenClaims struct {
	UserUUID string `json:"user_uuid"`
	jwt.StandardClaims
}

type refreshTokenClaims struct {
	UserUUID string `json:"user_uuid"`
	jwt.StandardClaims
}

func NewToken(userUUID string) (*Token, error) {
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims{
		userUUID,
		jwt.StandardClaims{
			Issuer:    "LYnLab/Auth",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().AddDate(0, 0, 7).Unix(),
		},
	}).SignedString(hmacSecret)
	if err != nil {
		return nil, err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims{
		userUUID,
		jwt.StandardClaims{
			Issuer:    "LYnLab/Auth",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().AddDate(0, 0, 28).Unix(),
		},
	}).SignedString(hmacSecret)
	if err != nil {
		return nil, err
	}

	return &Token{
		UserUUID:     userUUID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// Validate function returns UUID of token owner, if the token is valid.
func (token *Token) Validate() (string, error) {
	accessToken, err := jwt.ParseWithClaims(token.AccessToken, &accessTokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := accessToken.Claims.(*accessTokenClaims); ok && accessToken.Valid {
		return claims.UserUUID, nil
	} else {
		return "", fmt.Errorf("Invalid claims")
	}
}

func init() {
	hmacSecret = []byte(os.Getenv("LYNLAB_SECRET_KEY"))
}
