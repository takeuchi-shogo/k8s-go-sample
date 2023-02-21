package gateways

import jwt "github.com/dgrijalva/jwt-go"

type Jwt interface {
	CreateToken(userID int) string
	ParseToken(token string) (jwt.MapClaims, error)
}
