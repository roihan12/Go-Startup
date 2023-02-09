package auth

import (
	"github.com/golang-jwt/jwt/v4"
)

type Service interface {
	GenerateToken(userID int) (string, error)
}

var SecretKey = []byte("contohkEY123")

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SecretKey)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil

}
