package constants

import "github.com/golang-jwt/jwt"

// TODO: use an RSA encrypted key
var JwtKey = []byte("Nishant_secret_key")

var SignMethod = jwt.SigningMethodHS256
