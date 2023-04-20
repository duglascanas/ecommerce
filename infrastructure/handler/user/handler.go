package user

import (
	"github.com/duglascanas/ecommerce/domain/user"
	"github.com/duglascanas/ecommerce/model"
	"github.com/labstack/echo/v4"
)

type handler struct {
	useCase user.UseCase
}

func newHandler(uc user.UseCase) handler {
	return handler{useCase: uc}
}

func (h handler) Create(c echo.Context) error {
	m := model.User{}

	if err := c.Bind(&m); err != nil {
		return h.responser.BindFailed(err)
	}
}
