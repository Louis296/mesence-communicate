package jwt

import (
	"github.com/dgrijalva/jwt-go"
)

var Secret string

type Claims struct {
	Phone string
	Name  string
	jwt.StandardClaims
}

func GenerateToken(phone, name string) (string, error) {
	claims := Claims{
		Phone: phone,
		Name:  name,
		StandardClaims: jwt.StandardClaims{
			//ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			Subject: "mesence",
			Issuer:  "mesence_louis296",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(Secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
