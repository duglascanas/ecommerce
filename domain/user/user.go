package user

import (
	"github.com/duglascanas/ecommerce/model"
	"github.com/google/uuid"
)

type UseCase interface {
	Create(m *model.User) error
	GetbyID(ID uuid.UUID) (model.User, error)
	GetByEmail(email string) (model.User, error)
	GetAll() (model.Users, error)
}

type Storage interface {
	Create(m *model.User) error
	GetByID(ID uuid.UUID) (model.User, error)
	GetByEmail(email string) (model.User, error)
	GetAll() (model.Users, error)
}
