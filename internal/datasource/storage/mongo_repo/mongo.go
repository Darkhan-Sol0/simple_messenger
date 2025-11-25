package mongo_repo

import "go.mongodb.org/mongo-driver/v2/mongo"

type (
	mongo_repo struct {
		client *mongo.Client
	}

	Storage interface{}
)

func New(client *mongo.Client) Storage {
	return &mongo_repo{
		client: client,
	}
}
