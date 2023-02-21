package middleware

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/takeuchi-shogo/k8s-go-sample/config"
)

type Jwt struct {
	ApplicationName string
	TokenExpireAt   int
	SecretKey       string
}

type UnsignedResponse struct {
	Message interface{} `json:"message"`
}

func NewJwtAuth(c *config.Config) *Jwt {
	return &Jwt{
		ApplicationName: c.ApplicationName,
		TokenExpireAt:   c.TokenExpireAt,
		SecretKey:       c.SecretKey,
	}
}

// Create json web token
func (j *Jwt) CreateToken(userID int) string {
	log.Println("6", userID, j.ApplicationName, j.TokenExpireAt)
	claim := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * time.Duration(j.TokenExpireAt)).Unix(),
		Issuer:    j.ApplicationName,
		Audience:  strconv.Itoa(userID),
	}
	// ES256 には公開鍵と秘密鍵のペアが必要で、HS256 には秘密鍵のみが必要
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// Add Signature to Token
	tokenString, _ := token.SignedString([]byte(j.SecretKey))
	return tokenString
}

// Parse json web token
func (j *Jwt) ParseToken(jwtToken string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.SecretKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); !ok && token.Valid {
		return claims, nil
	} else {
		return claims, err
	}
}

func extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("bad header value given")
	}

	jwtToken := strings.Split(header, ".")
	if len(jwtToken) != 3 {
		return "", errors.New("incorrectly formatted authorization header")
	}

	return header, nil
}

// Check json web token in request header
func (j *Jwt) CheckJwtToken(c *gin.Context) {
	token, err := extractBearerToken(c.GetHeader("Authorization"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: err.Error(),
		})
		return
	}
	_, err = j.ParseToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, UnsignedResponse{
			Message: err.Error(),
		})
		return
	}
	return
}
