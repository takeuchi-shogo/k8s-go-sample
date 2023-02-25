package helpers

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

func SetHeader(ctx context.Context, token string) error {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return err
	}

	gc.Header("Authorization", "Bearer "+token)

	return nil
}

// context to gin.Context
func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value("GinContextKey")
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}
