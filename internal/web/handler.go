package web

import (
	"net/http"
	"simple_messenger/internal/dto"

	"github.com/labstack/echo/v4"
)

func (r *routing) Hello(ctx echo.Context) error {
	return ctx.JSON(200, map[string]string{"data": "hui"})
}

func (r *routing) NewChat(ctx echo.Context) error {
	var data dto.NewChatDTO
	err := ctx.Bind(&data)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	res, err := r.service.Chat().NewChat(ctx.Request().Context(), &data)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, map[string]any{"data": res})
}

func (r *routing) GetAllChats(ctx echo.Context) error {
	res, err := r.service.Chat().GetAllChats(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, map[string]any{"data": res})
}

func (r *routing) GetAllMyChats(ctx echo.Context) error {
	var data dto.UUIDUser
	err := ctx.Bind(&data)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	res, err := r.service.Chat().GetAllMyChats(ctx.Request().Context(), &data)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, map[string]any{"data": res})
}

func (r *routing) SendMessange(ctx echo.Context) error {
	var data dto.SendMessange
	err := ctx.Bind(&data)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	res, err := r.service.Messenger().SendMessange(ctx.Request().Context(), &data)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, map[string]any{"data": res})
}

func (r *routing) GetMessange(ctx echo.Context) error {
	var data dto.UUIDChat
	err := ctx.Bind(&data)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	res, err := r.service.Messenger().GetMessange(ctx.Request().Context(), &data)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, map[string]any{"data": res})
}
