package server

import (
	"simple_message/internal/config"
	"simple_message/internal/datasource/database"
	"simple_message/internal/service"
	"simple_message/internal/web"

	"github.com/labstack/echo/v4"
)

type (
	server struct {
		httpDriver *echo.Echo
		cfg        config.Config
		database   database.Database
		service    service.Service
		router     web.Routing
	}

	Server interface {
	}
)

func New() Server {
	return &server{
		httpDriver: echo.New(),
		cfg:        config.GetConfig(),
	}
}
