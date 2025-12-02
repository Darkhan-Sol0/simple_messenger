package mongo_repo

import (
	"context"
	"fmt"
	"simple_messenger/internal/dto"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type (
	mongo_repo struct {
		client *mongo.Client
	}

	Storage interface {
		SendMessange(ctx context.Context, data *dto.SendMessange) (bool, error)
		GetMessange(ctx context.Context, data *dto.UUIDChat) ([]dto.GetMessange, error)
	}
)

func New(client *mongo.Client) Storage {
	return &mongo_repo{
		client: client,
	}
}

func (m *mongo_repo) SendMessange(ctx context.Context, data *dto.SendMessange) (bool, error) {
	collection := m.client.Database("mongodb").Collection("messange")
	_, err := collection.InsertOne(ctx, *data)
	if err != nil {
		return false, fmt.Errorf("error: %v", err)
	}
	return true, nil
}

func (m *mongo_repo) GetMessange(ctx context.Context, data *dto.UUIDChat) ([]dto.GetMessange, error) {
	collection := m.client.Database("mongodb").Collection("messange")
	cur, err := collection.Find(ctx, *data)
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}
	defer cur.Close(ctx)
	var ms []dto.GetMessange
	for cur.Next(ctx) {
		var m dto.GetMessange
		err := cur.Decode(&m)
		if err != nil {
			return nil, fmt.Errorf("error: %v", err)
		}
		ms = append(ms, m)
	}
	return ms, nil
}
