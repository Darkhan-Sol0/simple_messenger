package postgres_repo

import (
	"context"
	"fmt"
	"simple_messenger/internal/dto"

	"github.com/Darkhan-Sol0/simple_qb"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	pg_repo struct {
		client *pgxpool.Pool
	}

	Storage interface {
		CreateNewChat(ctx context.Context, data *dto.NewChatDTO) (*dto.UUIDChat, error)
		GetAllChats(ctx context.Context) ([]dto.UUIDChat, error)
		GetAllMyChats(ctx context.Context, data *dto.UUIDUser) ([]dto.UUIDChat, error)
	}
)

var table = "chats"

func New(client *pgxpool.Pool) Storage {
	return &pg_repo{
		client: client,
	}
}

func (pg *pg_repo) CreateNewChat(ctx context.Context, data *dto.NewChatDTO) (*dto.UUIDChat, error) {
	qb := simple_qb.New(table, *data, nil)
	query, args := qb.Insert() //uuid_chat
	query = fmt.Sprintf("%s RETURNING uuid_chat", query)
	res := dto.UUIDChat{}
	err := pg.client.QueryRow(ctx, query, args...).Scan(&res.UUIDChat)
	if err != nil {
		return nil, fmt.Errorf("error: add new chat: %v", err.Error())
	}
	fmt.Println(res)
	return &res, nil
}

func (pg *pg_repo) GetAllChats(ctx context.Context) ([]dto.UUIDChat, error) {
	qb := simple_qb.New(table, dto.UUIDChat{}, nil)
	query, _ := qb.Select()
	fmt.Println(query)
	rows, err := pg.client.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	dataOut, err := pgx.CollectRows(rows, pgx.RowToStructByName[dto.UUIDChat])
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return dataOut, nil
}

func (pg *pg_repo) GetAllMyChats(ctx context.Context, data *dto.UUIDUser) ([]dto.UUIDChat, error) {
	filter := simple_qb.NewParam(simple_qb.NewNode("uuid_user_1", "eq", data.UUID), simple_qb.NewNodeOr("uuid_user_2", "eq", data.UUID))
	qb := simple_qb.New(table, dto.UUIDChat{}, filter)
	query, args := qb.Select()
	fmt.Println(query, args)
	rows, err := pg.client.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	dataOut, err := pgx.CollectRows(rows, pgx.RowToStructByName[dto.UUIDChat])
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return dataOut, nil
}
