-- +goose Up
-- +goose StatementBegin


CREATE TABLE group_members (
                               group_id UUID REFERENCES chat_groups(id) ON DELETE CASCADE,
                               user_id UUID REFERENCES users(id) ON DELETE CASCADE,
                               joined_at TIMESTAMPTZ DEFAULT NOW(),

                               PRIMARY KEY (group_id, user_id)
);

-- +goose StatementEnd
-- +goose Down

-- +goose StatementBegin
DROP TABLE group_members;
-- +goose StatementEnd
