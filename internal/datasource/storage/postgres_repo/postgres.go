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
	filter1 := simple_qb.NewFillter(simple_qb.NewNode("uuid_user_1", "eq", data.UUID))
	filter2 := simple_qb.NewFillter(simple_qb.NewNode("uuid_user_2", "eq", data.UUID))
	qb := simple_qb.New(table, dto.UUIDChat{}, filter1)
	query1, _ := qb.Select()
	qb = simple_qb.New(table, dto.UUIDChat{}, filter2)
	query2, _ := qb.Select()
	rows1, err := pg.client.Query(ctx, query1)
	if err != nil {
		return nil, err
	}
	rows2, err := pg.client.Query(ctx, query2)
	if err != nil {
		return nil, err
	}
	defer rows1.Close()
	if err := rows1.Err(); err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	defer rows2.Close()
	if err := rows2.Err(); err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	dataOut1, err := pgx.CollectRows(rows1, pgx.RowToStructByName[dto.UUIDChat])
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	dataOut2, err := pgx.CollectRows(rows1, pgx.RowToStructByName[dto.UUIDChat])
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	dataOut := append(dataOut1, dataOut2...)
	return dataOut, nil
}
