package common

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type MyCustomClaims struct {
	UserID string
	jwt.StandardClaims
}

func GetToken(secret, uid string) string {
	expireToken := time.Now().Add(time.Hour * 24 * 30).Unix()
	claims := MyCustomClaims{
		uid,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "sxiot",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte(secret))
	return signedToken
}

func GetUID(tokenStr, secret string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims.UserID, nil
	}

	return "", errors.New("非法操作")

}
