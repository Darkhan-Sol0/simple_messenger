-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    login TEXT,
    is_active BOOL,
    date_created TIMESTAMPZ,
    date_delete TIMESTAMPZ
);

CREATE TABLE IF NOT EXISTS chats (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    is_active BOOL,
    date_created TIMESTAMPZ,
    date_delete TIMESTAMPZ
);

CREATE TABLE IF NOT EXISTS user_subs_chats (
    id SERIAL PRIMARY KEY,
    user_uuid UUID,
    chat_uuid UUID,
    is_active BOOL,
    date_created TIMESTAMPZ
);

-- Индексация полей для ускорения фильтрации
CREATE INDEX idx_is_active_users ON users(is_active);
CREATE INDEX idx_user_subs_chats_chat ON user_subs_chats(chat_uuid);
CREATE INDEX idx_user_subs_chats_user ON user_subs_chats(user_uuid);
CREATE INDEX idx_is_active_subs_chats ON user_subs_chats(is_active); -- Здесь правильное название индекса
CREATE INDEX idx_is_active_chats ON chats(is_active);

-- Внешние ключи
ALTER TABLE user_subs_chats ADD CONSTRAINT fk_user FOREIGN KEY (user_uuid) REFERENCES users(id);
ALTER TABLE user_subs_chats ADD CONSTRAINT fk_chat FOREIGN KEY (chat_uuid) REFERENCES chats(id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_subs_chats;
DROP TABLE users;
DROP TABLE chats;
-- +goose StatementEnd