package routes

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/takeuchi-shogo/k8s-go-sample/config"
	"github.com/takeuchi-shogo/k8s-go-sample/graph"
	"github.com/takeuchi-shogo/k8s-go-sample/infrastructure/database"
	"github.com/takeuchi-shogo/k8s-go-sample/infrastructure/middleware"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/controllers"
)

type Routes struct {
	DB  *database.DB
	Gin *gin.Engine
}

func NewRoutes(c *config.Config, db *database.DB) *Routes {
	r := &Routes{
		DB:  db,
		Gin: gin.Default(),
	}

	r.setCors(middleware.NewCors(c))

	r.setRouting()

	return r
}

func (r *Routes) setCors(cors *middleware.Cors) {
	r.Gin.Use(cors.Config)
}

func (r *Routes) setRouting() {

	usersController := controllers.NewUsersController(r.DB)

	// srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	r.Gin.GET("/", playGround())
	r.Gin.POST("/query", graphqlHandler())

	r.Gin.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hellow world!!",
		})
	})

	v1 := r.Gin.Group("/v1")

	v1.POST("/users", func(c *gin.Context) { usersController.Post(c) })
	v1.GET("/users/:id", func(c *gin.Context) { usersController.Get(c) })
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	// Post only
	//
	// Since we are not using Playground, Get is not necessary.
	// h.AddTransport(transport.POST{})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playGround() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
