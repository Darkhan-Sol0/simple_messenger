package web

import (
	"simple_message/internal/service"

	"github.com/labstack/echo/v4"
)

type (
	routing struct {
		service service.Service
	}

	Routing interface {
		RegisterRoutes(e *echo.Echo)
	}
)

func New(service service.Service) Routing {
	return &routing{
		service: service,
	}
}

func (r *routing) RegisterRoutes(e *echo.Echo) {

}
