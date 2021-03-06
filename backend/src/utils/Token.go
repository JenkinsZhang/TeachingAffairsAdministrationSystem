package utils

import (
	"errors"
	"fmt"
	"taas/models"

	"github.com/dgrijalva/jwt-go"
)

func NewToken(payload jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString([]byte(models.SECRET))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func CheckToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(models.SECRET), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if err := UserExists(claims["id"].(string)); err == nil {
			return claims, nil
		}
		return nil, errors.New("invalid token")
	}
	return nil, err

}

func UserExists(id string) error {
	var identity string
	err := Db.QueryRow("select identity from User where id = ?", id).Scan(&identity)
	return err
}
