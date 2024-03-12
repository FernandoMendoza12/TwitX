package jwt

import (
	"errors"
	"strings"

	"github.com/FernandoMendoza12/TwitX/models"
	jwt "github.com/golang-jwt/jwt/v5"
)

var Email string
var IDUser string

func ProcessToken(tk string, JWTSign string) (*models.Claim, bool, string, error) {
	myKey := []byte(JWTSign)

	var claims models.Claim

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return &claims, false, "", errors.New("incorrect token format")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, &claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})

	if err == nil {
		//Routine check agains the DB
	}

	if !tkn.Valid {
		return &claims, false, string(""), errors.New("invaled Token")
	}

	return &claims, false, string(""), err
}
