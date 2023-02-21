package gateways

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type JwtGateway struct {
	Jwt Jwt
}

func (j *JwtGateway) CreateToken(userID int) string {
	return j.Jwt.CreateToken(userID)
}

func (j *JwtGateway) ParseToken(token string) (jwt.MapClaims, error) {
	return j.Jwt.ParseToken(token)
}
