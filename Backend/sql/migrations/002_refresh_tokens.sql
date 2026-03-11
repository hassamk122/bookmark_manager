-- +goose Up
CREATE TABLE IF NOT EXISTS refresh_tokens(
    token_id SERIAL CONSTRAINT rt_key PRIMARY KEY,
    token varchar(255) NOT NULL UNIQUE,
    user_id int REFERENCES users(users_id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE refresh_tokens;