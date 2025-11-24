package util

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JwtClaims struct {
	UserId   int64  `json:"userId"`
	UserName string `json:"userName"`
	jwt.RegisteredClaims
}

// ParseToken 解析JWT token
func ParseToken(tokenString string, secret string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// GenerateToken generates a JWT token for the given user information
// Parameters:
//   - userId: the unique identifier of the user
//   - username: the name of the user
//   - secret: the secret key used to sign the token
//   - expireDuration: the duration after which the token will expire
//
// Returns:
//   - string: the generated JWT token string
//   - error: any error that occurred during token generation
func GenerateToken(userId int64, username string, secret string, expireDuration time.Duration) (string, error) {
	// Set the current time and calculate expiration time
	now := time.Now()
	expireAt := now.Add(expireDuration).Unix()

	// Create JWT claims with user information and token metadata
	claims := &JwtClaims{
		UserId:   userId,
		UserName: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(expireAt, 0)),
			IssuedAt:  jwt.NewNumericDate(now),
			Issuer:    "your-issuer",
		},
	}

	// Generate and sign the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
