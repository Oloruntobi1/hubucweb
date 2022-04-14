package repository

import (
	"fmt"

	"github.com/Oloruntobi1/hubucweb/internal/models"
	"github.com/google/uuid"
)

var (
	ErrUserNotFound      = fmt.Errorf("user not found")
	ErrUserAlreadyEXists = fmt.Errorf("user already exists")
)

type Reader interface {
	GetUser(id uuid.UUID) (*models.User, error)
}

type Creater interface {
	CreateUser(uuid.UUID, *models.User) (*models.User, error)
}

type Repository interface {
	Reader
	Creater
}

func GetStorage(s string) Repository {
	switch {
	case s == "MYSQL":
		// return MYSQL implementation
	default:
		inMemory := NewInMemory()
		return inMemory
	}

	return nil
}
