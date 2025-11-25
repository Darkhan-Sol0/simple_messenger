package chat

import (
	"context"
	"simple_messenger/internal/datasource/database"
	"simple_messenger/internal/datasource/storage"
	"simple_messenger/internal/dto"
)

type (
	chat struct {
		storage storage.Repo
	}

	Chat interface {
		NewChat(ctx context.Context, data *dto.NewChatDTO) (*dto.UUIDChat, error)
		GetAllChats(ctx context.Context) ([]dto.UUIDChat, error)
		GetAllMyChats(ctx context.Context, data *dto.UUIDUser) ([]dto.UUIDChat, error)
	}
)

func New(db database.Database) Chat {
	return &chat{
		storage: storage.New(db),
	}
}

func (c *chat) NewChat(ctx context.Context, data *dto.NewChatDTO) (*dto.UUIDChat, error) {
	res, err := c.storage.Postgres().CreateNewChat(ctx, data)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *chat) GetAllChats(ctx context.Context) ([]dto.UUIDChat, error) {
	res, err := c.storage.Postgres().GetAllChats(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *chat) GetAllMyChats(ctx context.Context, data *dto.UUIDUser) ([]dto.UUIDChat, error) {
	res, err := c.storage.Postgres().GetAllMyChats(ctx, data)
	if err != nil {
		return nil, err
	}
	return res, nil
}
