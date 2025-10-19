package auth

import (
	"strconv"
	"time"

	"github.com/Nutan-Kum12/Ecom/config"
	"github.com/golang-jwt/jwt"
)

func CreateJWT(userID int, secretKey []byte) (string, error) {
	// Implementation for creating JWT
	expiration := time.Second * time.Duration(config.Envs.JWTExpirationInSeconds)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": strconv.Itoa(userID),
		"exp":     time.Now().Add(expiration).Unix(),
	})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
