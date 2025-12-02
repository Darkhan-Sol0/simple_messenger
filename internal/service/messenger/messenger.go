package messenger

import (
	"context"
	"simple_messenger/internal/datasource/database"
	"simple_messenger/internal/datasource/storage"
	"simple_messenger/internal/dto"
)

type (
	messenger struct {
		storage storage.Repo
	}

	Messenger interface {
		SendMessange(ctx context.Context, data *dto.SendMessange) (bool, error)
		GetMessange(ctx context.Context, data *dto.UUIDChat) ([]dto.GetMessange, error)
	}
)

func New(db database.Database) Messenger {
	return &messenger{
		storage: storage.New(db),
	}
}

func (m *messenger) SendMessange(ctx context.Context, data *dto.SendMessange) (bool, error) {
	res, err := m.storage.Mongo().SendMessange(ctx, data)
	if err != nil {
		return false, err
	}
	return res, nil
}

func (m *messenger) GetMessange(ctx context.Context, data *dto.UUIDChat) ([]dto.GetMessange, error) {
	res, err := m.storage.Mongo().GetMessange(ctx, data)
	if err != nil {
		return nil, err
	}
	return res, nil
}
