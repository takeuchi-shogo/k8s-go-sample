package server

type Server struct {
	Port string
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run(port string) {
	s.Run(port)
}
