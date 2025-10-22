package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"fmt"
	"time"
)

type UserClaims struct {
	UserID int `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(userID int, role string, secretKey []byte) (string, error) {

	//expirationTime := time.Now().Add(24 * time.Hour)

	shortExpirationTime := time.Now().Add(15 * time.Minute)

	claims := &UserClaims{

		UserID: userID,
		Role: role,

		RegisteredClaims: jwt.RegisteredClaims{

			ExpiresAt: jwt.NewNumericDate(shortExpirationTime),
			IssuedAt: jwt.NewNumericDate(time.Now()),

		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", fmt.Errorf("could not sign token: %w", err)
	}

	return tokenString, nil
}