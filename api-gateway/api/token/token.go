package tokens

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)
func VerifyToken(tokenString string) error {
	secretKey := []byte("secret")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
