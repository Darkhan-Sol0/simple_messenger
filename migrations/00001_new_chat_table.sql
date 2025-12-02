-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS chats (
    chat_uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_uuid_1 TEXT,
    user_uuid_2 TEXT,
    date_created TIMESTAMPZ
);

CREATE INDEX IF NOT EXISTS idx_user_uuid_1 ON chats(user_uuid_1);
CREATE INDEX IF NOT EXISTS idx_user_uuid_2 ON chats(user_uuid_2);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE chats;
-- +goose StatementEnd