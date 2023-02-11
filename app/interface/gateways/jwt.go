package gateways

import jwt "github.com/dgrijalva/jwt-go"

type Jwt interface {
	ParseToken(token string) (jwt.MapClaims, error)
}
