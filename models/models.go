package models

import (
	"github.com/golang-jwt/jwt"
	"golang.org/x/oauth2"
)

type Response struct {
	ResponseCode    string      `json:"responseCode"`
	ResponseMessage string      `json:"responseMessage"`
	AdditionalInfo  interface{} `json:"additionalInfo"`
}

type OidcData struct {
	AuthCode  string
	AuthToken *oauth2.Token
	UserData  jwt.Claims
}
