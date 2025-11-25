package mongodb

import (
	"context"
	"fmt"
	"log"
	"simple_messenger/internal/config"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func ConnectMongo(ctx context.Context, cfg config.Config) (pool *mongo.Client, err error) {
	dns := fmt.Sprintf("mongodb://%s:%s",
		cfg.GetMongoHost(),
		cfg.GetMongoPort(),
	)
	clientOptions := options.Client().ApplyURI(dns)
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return nil, fmt.Errorf("не удалось подключиться к MongoDB: %v", err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, fmt.Errorf("ошибка Ping в MongoDB: %v", err)
	}
	log.Println("Подключился к MongoDB")
	return client, nil
}
