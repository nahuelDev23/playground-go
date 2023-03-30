package services

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/nahueldev23/playground/models"
)

func GenerateToken(data *models.Login) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"name": data.Name,
	})
	signedToken, err := token.SignedString([]byte("joderestoesarte"))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}
