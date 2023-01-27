package server

import "github.com/takeuchi-shogo/k8s-go-sample/infrastructure/routes"

type Server struct {
	Routes *routes.Routes
	Port   string
}

func NewServer(routes *routes.Routes) *Server {
	return &Server{
		Routes: routes,
	}
}

func (s *Server) Run(port string) {
	s.Routes.Gin.Run(port)
}
