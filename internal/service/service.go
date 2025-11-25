package service

import (
	"simple_messenger/internal/datasource/database"
	"simple_messenger/internal/service/chat"
	"simple_messenger/internal/service/messenger"
)

type (
	service struct {
		chat      chat.Chat
		messenger messenger.Messenger
	}

	Service interface {
		Chat() chat.Chat
		Messenger() messenger.Messenger
	}
)

func New(db database.Database) Service {
	return &service{
		chat:      chat.New(db),
		messenger: messenger.New(db),
	}
}

func (s *service) Chat() chat.Chat {
	return s.chat
}

func (s *service) Messenger() messenger.Messenger {
	return s.messenger
}
