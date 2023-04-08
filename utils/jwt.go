package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"time"
)

var stSigningKey = []byte(viper.GetString("jwt.sigingKey"))

type JwtCustClaims struct {
	ID   int
	Name string
	jwt.RegisteredClaims
}

func GenerateToken(id int, name string) (string, error) {
	jwtCustClaims := JwtCustClaims{
		ID:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "nbgo",
			Subject:   "Token",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(viper.GetDuration("jwt.expiresAt") * time.Minute)),
			NotBefore: nil,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtCustClaims)
	return token.SignedString(stSigningKey)
}

func ParseToken(tokenStr string) (JwtCustClaims, error) {
	jwtCustClaims := JwtCustClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, &jwtCustClaims, func(token *jwt.Token) (interface{}, error) {
		return stSigningKey, nil
	})
	if err == nil && !token.Valid {
		err = errors.New("Invalid Token")
	}

	return jwtCustClaims, err
}

func IsTokenValid(tokenStr string) bool {
	_, err := ParseToken(tokenStr)
	if err != nil {
		return false
	}

	return true
}
