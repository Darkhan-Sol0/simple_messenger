package dto

import "time"

type (
	NewChatDTO struct {
		UUIDMainer string `json:"user_uuid_1" db:"user_uuid_1"`
		UUIDUser   string `json:"user_uuid_2" db:"user_uuid_2"`
	}

	UUIDChat struct {
		UUIDChat string `json:"chat_uuid" db:"chat_uuid" bson:"chat_uuid"`
	}

	SendMessange struct {
		UUIDChat string        `json:"chat_uuid" bson:"chat_uuid"`
		Text     string        `json:"text" bson:"text"`
		Date     time.Duration `json:"date" bson:"date"`
	}

	GetMessange struct {
		Text string        `json:"text" bson:"text"`
		Date time.Duration `json:"date" bson:"date"`
	}

	UUIDUser struct {
		UUID string `json:"uuid"`
	}
)
