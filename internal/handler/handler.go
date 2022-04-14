package handler

import (
	repo "github.com/Oloruntobi1/hubucweb/internal/repository"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  repo.Repository
	router *gin.Engine
}

func NewServer(store repo.Repository)( *Server, error) {
	s := &Server{
		store: store,
	}
	s.setupRouter()
	return s, nil

}

func (s *Server) setupRouter() {
	r := gin.Default()

	r.POST("/api/v1/user", s.CreateUser)
	s.router = r
}

func (s *Server) Start(port string) error {
	return s.router.Run(":" + port)
}

// type

type GoodResponse struct {
	Success bool
	Message string
	Data    interface{}
}

type BadResponse struct {
	Message string
	Err     error
}
