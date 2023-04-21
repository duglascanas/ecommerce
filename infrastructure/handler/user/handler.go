package user

import (
	"errors"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/duglascanas/ecommerce/domain/user"
	"github.com/duglascanas/ecommerce/model"
	//"github.com/duglascanas/ecommerce/infrastructure/handler/response"
)

type handler struct {
	useCase   user.UseCase
	responser response.API
}

func newHandler(uc user.UseCase) handler {
	return handler{useCase: uc}
}

func (h handler) Create(c echo.Context) error {
	m := model.User{}

	if err := c.Bind(&m); err != nil {
		return h.responser.BindFailed(err)
	}

	if err := h.useCase.Create(&m); err != nil {
		return h.responser.Error(c, "useCreate()", err)
	}

	return c.JSON(h.responser.Create(m))

}

// Myself returns the data from my profile
func (h handler) MySelf(c echo.Context) error {
	ID, ok := c.Get("userID").(uuid.UUID)
	if !ok {
		return &h.responser.Error(c, "c.Get().(uuid.UUID)", errors.New("couldn´t parse the ID"))
	}

	u, err := h.useCase.GetByID(ID)
	if err != nil {
		return h.responser.Error(c, "useCase.GetWhere()", err)
	}

	return c.JSON(h.responser.OK(u))
}

func (h handler) GetAll(c echo.Context) error {
	users, err := h.useCase.GetAll()
	if err != nil {
		return h.responser.Error(c, "useCase.GetAll()", err)
	}

	return c.JSON(h.responser.OK(users))
}
