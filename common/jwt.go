package common

import (
	"ginvue/model"
	_"go/token"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)	// token有效时间
	claims := &Claims {
		UserId : user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),	// 有效时间
			IssuedAt: time.Now().Unix(),	// 发放时间
			Issuer: "ginvue", 		// 发放者
			Subject: "user token", 	// 主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}


func ParseToken(tokenString string) (*jwt.Token, *Claims, error){
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
		
	})

	return token, claims, err
}
