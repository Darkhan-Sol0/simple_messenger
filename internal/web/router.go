package web

import (
	"simple_messenger/internal/service"

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
	e.GET("/health", r.Hello)
	e.POST("/chat", r.NewChat)
	e.GET("/chats", r.GetAllChats)
	e.GET("/mychats", r.GetAllMyChats)

	e.POST("/send", r.SendMessange)
	e.GET("/get", r.GetMessange)
}
