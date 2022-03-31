package services

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type authCustomClaims struct {
	EmailAddress string
	jwt.StandardClaims
}

func (service DefaultService) GenerateToken(email string) (string, error) {
	claims := &authCustomClaims{
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(service.Config.Secret))
	if err != nil {
		return "", err
	}
	return t, nil
}

func (service DefaultService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, errors.New("invalid token")
		}
		return []byte(service.Config.Secret), nil
	})

}