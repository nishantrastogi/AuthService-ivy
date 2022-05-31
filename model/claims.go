package model

import "github.com/golang-jwt/jwt"

// Adding the fields that should be a part of the payload in the token.
type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	// standard claims is added as an embedded type to provide fields like expiry time.
	jwt.StandardClaims
}

func NewClaims(username string, role string, standardClaims jwt.StandardClaims) *Claims {
	r := &Claims{username, role, standardClaims}
	return r
}
