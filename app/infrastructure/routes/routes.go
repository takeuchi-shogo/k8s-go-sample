package routes

import (
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/takeuchi-shogo/k8s-go-sample/config"
	"github.com/takeuchi-shogo/k8s-go-sample/graphql"
	"github.com/takeuchi-shogo/k8s-go-sample/infrastructure/database"
	"github.com/takeuchi-shogo/k8s-go-sample/infrastructure/middleware"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/controllers"
)

type Routes struct {
	DB  *database.DB
	Gin *gin.Engine
	Jwt *middleware.Jwt
}

func NewRoutes(c *config.Config, db *database.DB, jwt *middleware.Jwt) *Routes {
	r := &Routes{
		DB:  db,
		Gin: gin.Default(),
		Jwt: jwt,
	}

	r.setCors(middleware.NewCors(c))

	r.setRouting()

	return r
}

func (r *Routes) setCors(cors *middleware.Cors) {
	r.Gin.Use(cors.Config)
}

func (r *Routes) setRouting() {

	accountsController := controllers.NewAccountsController(controllers.AccountsControllerProvider{DB: r.DB, Jwt: r.Jwt})
	blocksController := controllers.NewBlocksController(controllers.BlocksControllerProvider{DB: r.DB, Jwt: r.Jwt})
	reportsController := controllers.NewReportsController(controllers.ReportsControllerProvider{DB: r.DB, Jwt: r.Jwt})
	userProfilesController := controllers.NewUserProfilesController(controllers.UserProfilesControllerProvider{DB: r.DB, Jwt: r.Jwt})
	usersController := controllers.NewUsersController(controllers.UsersControllerProvider{DB: r.DB, Jwt: r.Jwt})

	// srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	// GraphQL 用
	r.Gin.Use(middleware.GinContextToContextMiddleware())
	r.Gin.GET("/", playGround())
	r.Gin.POST("/query", r.graphqlHandler())

	r.Gin.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello world!!",
		})
	})

	// REST API用
	v1 := r.Gin.Group("/v1")

	// accounts
	v1.GET("/accounts", func(ctx *gin.Context) { accountsController.Get(ctx) })
	v1.POST("/accounts", func(ctx *gin.Context) { accountsController.Post(ctx) })

	// blocks
	v1.GET("/blocks/:id", func(ctx *gin.Context) { blocksController.Get(ctx) })
	v1.POST("/blocks", func(ctx *gin.Context) { blocksController.Post(ctx) })

	// reports
	v1.GET("/reports/:id", func(ctx *gin.Context) { reportsController.Get(ctx) })
	v1.POST("/reports", func(ctx *gin.Context) { reportsController.Post(ctx) })

	// userProfiles
	v1.GET("/userProfiles/:id", func(ctx *gin.Context) { userProfilesController.Get(ctx) })
	v1.POST("/userProfiles", func(ctx *gin.Context) { userProfilesController.Post(ctx) })

	// users
	v1.POST("/users", func(ctx *gin.Context) { usersController.Post(ctx) })
	v1.GET("/users/:id", func(ctx *gin.Context) { usersController.Get(ctx) })
}

// Defining the Graphql handler
func (r *Routes) graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{Resolvers: &graphql.Resolver{
		DB:  r.DB,
		Jwt: r.Jwt,
	}}))

	// Post only
	//
	// Since we are not using Playground, Get is not necessary.
	// h.AddTransport(transport.POST{})

	return func(c *gin.Context) {
		fmt.Println(c.Request.Header.Get("Authorization"))
		// Create JWT
		// jwtToken := "testJwtToken"

		// レスポンスヘッダーにJWTをセットする
		// c.Header("Authorization", "Bearer "+jwtToken)

		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playGround() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
