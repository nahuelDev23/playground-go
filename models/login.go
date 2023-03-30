package models

import "github.com/golang-jwt/jwt/v4"

type Login struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Claim struct {
	Name string `json:"name"`
	jwt.StandardClaims
}
