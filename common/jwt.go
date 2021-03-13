package common

import (
	"ginEssential/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func RealeaseToken(user model.User) (string, error) {
	expirationtime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationtime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "Daddy_wilen",
			Subject:   "user token",
		},
	}

	//将claims封装为完整的token:主要是加了头部和加密方法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//得到真正的完整的token
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return token, claims, err
}
