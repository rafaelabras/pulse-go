-- +goose Up
-- +goose StatementBegin

CREATE TABLE chat_groups (
                             id UUID PRIMARY KEY,
                             name VARCHAR(100) NOT NULL,
                             description TEXT,
                             created_by UUID REFERENCES users(id) ON DELETE SET NULL,
                             created_at TIMESTAMPTZ DEFAULT NOW()
);

-- +goose StatementEnd
-- +goose Down

-- +goose StatementBegin
DROP TABLE chat_groups;
-- +goose StatementEnd
