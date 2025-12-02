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
	query, args := simple_qb.New(table).Insert(data).Returning("uuid_chat").Generate()
	res := dto.UUIDChat{}
	err := pg.client.QueryRow(ctx, query, args...).Scan(&res.UUIDChat)
	if err != nil {
		return nil, fmt.Errorf("error: add new chat: %v", err.Error())
	}
	fmt.Println(res)
	return &res, nil
}

func (pg *pg_repo) GetAllChats(ctx context.Context) ([]dto.UUIDChat, error) {
	query, _ := simple_qb.New(table).Select(dto.UUIDChat{}).Generate()
	fmt.Println(query)
	rows, err := pg.client.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}
	dataOut, err := pgx.CollectRows(rows, pgx.RowToStructByName[dto.UUIDChat])
	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}
	return dataOut, nil
}

func (pg *pg_repo) GetAllMyChats(ctx context.Context, data *dto.UUIDUser) ([]dto.UUIDChat, error) {
	query, args := simple_qb.New(table).Select(dto.UUIDChat{}).Params(simple_qb.NewParam("uuid_user_1", "eq", data.UUID), simple_qb.NewOrParam("uuid_user_2", "eq", data.UUID)).Generate()
	rows, err := pg.client.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}
	dataOut, err := pgx.CollectRows(rows, pgx.RowToStructByName[dto.UUIDChat])
	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}
	return dataOut, nil
}
