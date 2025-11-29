-- +goose Up
-- +goose StatementBegin

CREATE TABLE users (
                       id UUID PRIMARY KEY,
                       username VARCHAR(50) UNIQUE NOT NULL,
                       bio TEXT NOT NULL,
                       email VARCHAR(120) UNIQUE NOT NULL,
                       password_hash TEXT NOT NULL,
                       created_at TIMESTAMPTZ DEFAULT NOW(),
                       updated_at TIMESTAMPTZ DEFAULT NOW()
);
-- +goose StatementEnd
-- +goose Down

-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd

