package token

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type signedDetails struct {
	Email      string `json:"email"`
	First_Name string
	Last_Name  string
	Uid        string
	jwt.StandardClaims
}
