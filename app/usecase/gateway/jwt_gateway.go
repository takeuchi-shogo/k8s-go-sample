package gateway

import jwt "github.com/dgrijalva/jwt-go"

type JwtGateway interface {
	CreateToken(userID int) string
	ParseToken(token string) (jwt.MapClaims, error)
}
