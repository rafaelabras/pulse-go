-- +goose Up
-- +goose StatementBegin

CREATE TABLE messages (
                          id UUID PRIMARY KEY,
                          sender_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                          receiver_id UUID REFERENCES users(id) ON DELETE CASCADE,
                          group_id UUID REFERENCES chat_groups(id) ON DELETE CASCADE,
                          content TEXT,
                          content_type VARCHAR(20) DEFAULT 'text',  -- text, image, emoji, file, etc.
                          created_at TIMESTAMPTZ DEFAULT NOW()
                          CONSTRAINT valid_sent_message CHECK (
                              (receiver_id IS NOT NULL AND group_id IS NULL)
                              OR
                              (receiver_id IS NULL AND group_id IS NOT NULL)
                              )
);

-- +goose StatementEnd
-- +goose Down

-- +goose StatementBegin
DROP TABLE messages;
-- +goose StatementEnd
