-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS chats (
    uuid_chat UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    uuid_user_1 TEXT,
    uuid_user_2 TEXT
);

CREATE INDEX IF NOT EXISTS idx_uuid_user_1 ON chats(uuid_user_1);
CREATE INDEX IF NOT EXISTS idx_uuid_user_2 ON chats(uuid_user_2);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE chats;
-- +goose StatementEnd