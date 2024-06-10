package helper

import (
	"time"

	"github.com/cairomedeiros/todo-list/config"
	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(config.SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
