package repository

import (
	"sync"

	"github.com/Oloruntobi1/hubucweb/internal/models"
	"github.com/google/uuid"
)

type InMemory struct {
	//lock
	sync.Mutex
	UserMap map[uuid.UUID]*models.User
}

func NewInMemory() *InMemory {
	return &InMemory{
		UserMap: make(map[uuid.UUID]*models.User),
	}
}

//var _ Repository = (*InMemory)(nil)

func (i *InMemory) GetUser(id uuid.UUID) (*models.User, error) {
	v, ok := i.UserMap[id]
	if ok {
		return v, nil
	}

	return nil, ErrUserNotFound
}

func (i *InMemory) CreateUser(uuid uuid.UUID, user *models.User) (*models.User, error) {
	i.Lock()
	v, ok := i.UserMap[uuid]
	i.Unlock()
	// check email already exists
	if ok && v.Email == user.Email {
		return nil, ErrUserAlreadyEXists
	}

	i.Lock()
	i.UserMap[uuid] = user
	i.Unlock()

	return user, nil

}
