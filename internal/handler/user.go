package handler

import (
	"net/http"

	"github.com/Oloruntobi1/hubucweb/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (s *Server) CreateUser(ctx *gin.Context) {
	var req CreateUserRequest

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	// build the user model
	userID := uuid.New()

	u := &models.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}

	id, err := s.store.CreateUser(userID, u)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusCreated, id)
}
