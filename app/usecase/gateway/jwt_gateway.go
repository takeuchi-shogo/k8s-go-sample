package gateway

import jwt "github.com/dgrijalva/jwt-go"

type JwtGateway interface {
	ParseToken(token string) (jwt.MapClaims, error)
}
