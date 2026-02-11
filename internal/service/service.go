package service

import (
	"simple_message/internal/datasource/database"
	cuc "simple_message/internal/service/chats/usecase"
	uuc "simple_message/internal/service/users/usecase"
)

type (
	service struct {
		user uuc.UserUsecase
		chat cuc.ChatUsecase
	}

	Service interface {
		User() uuc.UserUsecase
		Chat() cuc.ChatUsecase
	}
)

func New(db database.Database) Service {
	return &service{
		user: uuc.New(),
		chat: cuc.New(),
	}
}

func (s *service) User() uuc.UserUsecase {
	return s.user
}

func (s *service) Chat() cuc.ChatUsecase {
	return s.chat
}
