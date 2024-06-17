package api

import "github.com/gin-gonic/gin"

type Server struct {
	engine *gin.Engine
	host   string
}

func (s *Server) Run() {
	s.setup
}
