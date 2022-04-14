package handler

import (
	"github.com/Oloruntobi1/hubucweb/internal/middleware"
	repo "github.com/Oloruntobi1/hubucweb/internal/repository"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  repo.Repository
	router *gin.Engine
}

func NewServer(store repo.Repository) (*Server, error) {
	s := &Server{
		store: store,
	}
	s.setupRouter()
	return s, nil

}

func (s *Server) setupRouter() {
	r := gin.Default()

	r.Use(middleware.LoggerToFile())
	r.POST("/user", s.CreateUser)
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

func errorResponse(err error) gin.H {
	return gin.H{
		"success": false,
		"message": "",
		"error":   err.Error(),
	}
}
