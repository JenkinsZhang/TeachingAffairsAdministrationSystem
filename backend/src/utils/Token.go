package utils

import (
	"database/sql"
	"errors"
	"fmt"
	"models"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func NewToken(payload jwt.MapClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString([]byte(models.SECRET))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Signed String: %v\n", err)
		os.Exit(-1)
	}
	return tokenString
}

func CheckToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(models.SECRET), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if ok := UserExists(claims["id"].(string)); ok {
			return claims, nil
		}
		return nil, errors.New("invalid token")
	}
	return nil, err

}

func UserExists(id string) bool {
	var identity string
	err := Db.QueryRow("select identity from User where id = ?", id).Scan(&identity)
	return err != sql.ErrNoRows
}
