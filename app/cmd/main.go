package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/takeuchi-shogo/k8s-go-sample/config"
)

func main() {

	config := config.NewConfig(".")
	fmt.Println(config)

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hellow world!!",
		})
	})

	r.Run(":8080")
}
