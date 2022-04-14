package repository

import (
	"github.com/Oloruntobi1/hubucweb/internal/models"
	"github.com/google/uuid"
)

type InMemory struct {
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

func (i *InMemory) CreateUser(uuid uuid.UUID, user *models.User) (uuid.UUID, error) {
	_, ok := i.UserMap[uuid]
	if ok {
		return uuid, nil
	}
	i.UserMap[uuid] = user

	return uuid, nil

}
