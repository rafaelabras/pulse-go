-- +goose Up
-- +goose StatementBegin


CREATE INDEX idx_users_username ON users(username);
CREATE INDEX idx_group_members_user ON group_members(user_id);
CREATE INDEX idx_messages_sender ON messages(sender_id);
CREATE INDEX idx_messages_receiver ON messages(receiver_id);
CREATE INDEX idx_messages_group ON messages(group_id);
CREATE INDEX idx_messages_created_at ON messages(created_at);

-- +goose StatementEnd
-- +goose Down

-- +goose StatementBegin
DROP INDEX idx_users_username;
DROP INDEX idx_group_members_user;
DROP INDEX idx_messages_sender;
DROP INDEX idx_messages_receiver;
DROP INDEX idx_messages_group;
DROP INDEX idx_messages_created_at;

-- +goose StatementEnd


