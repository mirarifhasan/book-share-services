package services

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

var secretKey string
var issuer string

type jwtCustomClaims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

func init() {
	secretKey = os.Getenv("JWT_SECRET")
	issuer = os.Getenv("JWT_ISSUER")
}

func GenerateToken(UserID int) string {
	claims := &jwtCustomClaims{
		UserID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(secretKey), nil
	})

}
