package handler

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"tkbai/config"
)

func GenerateJwtString(claims jwt.MapClaims) (jwtString string) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := []byte(config.JwtKey)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		config.LogErr(err, "Error signing token")
		return jwtString
	}

	return signedToken
}

func ParseJwtString(tokenString, claim string) (tokenPayload interface{}, err error) {
	secret := []byte(config.JwtKey)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		config.LogErr(err, "Error parsing token")
		return tokenPayload, err
	}

	if !token.Valid {
		err = errors.New("invalid token")
		config.LogErr(err, "Token is not valid")
		return tokenPayload, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		tokenPayload := claims[claim].(string)

		return tokenPayload, err
	} else {
		err = errors.New("error reading claims")
		config.LogErr(err, "Claim error")
		return tokenPayload, err
	}
}
