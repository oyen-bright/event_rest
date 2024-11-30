package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const SECRET_KEY = "THISISMYSECRETKEY"

func GenerateToke(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Duration(time.Hour * 2)).Unix(),
	})
	return token.SignedString([]byte(SECRET_KEY))

}

func ValidateToken(token string) (int64, error) {

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return 0, errors.New("unable to parse token")
	}
	if !parsedToken.Valid {
		return 0, errors.New("invalid toke")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invallid token claims")
	}
	// email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))

	return userId, nil
}
