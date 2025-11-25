package dto

import "time"

type (
	NewChatDTO struct {
		UUIDMainer string `json:"user_uuid1" db:"uuid_user_1"`
		UUIDUser   string `json:"user_uuid2" db:"uuid_user_2"`
	}

	UUIDChat struct {
		UUIDChat string `json:"chat_uuid" db:"chat_uuid"`
	}

	SendMessange struct {
		UUIDChat string        `json:"chat_uuid"`
		Text     string        `json:"text"`
		Date     time.Duration `json:"date"`
	}

	GetMessange struct {
		Text string        `json:"text"`
		Date time.Duration `json:"date"`
	}

	UUIDUser struct {
		UUID string `json:"uuid"`
	}
)
