package gateways

import jwt "github.com/dgrijalva/jwt-go"

type JwtGateway struct {
	Jwt Jwt
}

func (j *JwtGateway) ParseToken(token string) (jwt.MapClaims, error) {
	return j.Jwt.ParseToken(token)
}
