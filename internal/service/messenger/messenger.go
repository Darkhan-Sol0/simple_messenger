package messenger

import (
	"simple_messenger/internal/datasource/database"
	"simple_messenger/internal/datasource/storage"
)

type (
	messenger struct {
		storage storage.Repo
	}

	Messenger interface{}
)

func New(db database.Database) Messenger {
	return &messenger{
		storage: storage.New(db),
	}
}
