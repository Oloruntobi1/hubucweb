package handler

import (
	"net/http"

	"github.com/Oloruntobi1/hubucweb/internal/models"
	"github.com/Oloruntobi1/hubucweb/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=6,max=120"`
}

// We need this so we dont return the whole user model which might contain
// the user id which is bad to return to the user
type CreateUserResponse struct {
	Email string `json:"email"`
}

func (s *Server) CreateUser(ctx *gin.Context) {
	var req CreateUserRequest

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// build the user model
	userID := uuid.New()
	// hash password
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	u := &models.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		HashedPassword: hashedPassword,
	}

	user, err := s.store.CreateUser(userID, u)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ur := CreateUserResponse{
		Email: user.Email,
	}
	response := util.BuildResponseEntity(true, "Request Successful", gin.H{
		"data": ur,
	})
	ctx.JSON(http.StatusCreated, response)
}
