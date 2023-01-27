package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/takeuchi-shogo/k8s-go-sample/infrastructure/database"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/controllers"
)

type Routes struct {
	DB  *database.DB
	Gin *gin.Engine
}

func NewRoutes(db *database.DB) *Routes {
	r := &Routes{
		DB:  db,
		Gin: gin.Default(),
	}

	r.setRouting()

	return r
}

func (r *Routes) setRouting() {

	usersController := controllers.NewUsersController(r.DB)

	r.Gin.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hellow world!!",
		})
	})

	v1 := r.Gin.Group("/v1")

	v1.POST("/users", func(c *gin.Context) { usersController.Post(c) })
	v1.GET("/users/:id", func(c *gin.Context) { usersController.Get(c) })
}
