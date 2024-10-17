-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
ADD COLUMN email VARCHAR(255) NOT NULL;

ALTER TABLE users
ADD CONSTRAINT users_email_unique UNIQUE (email),
ADD CONSTRAINT users_username_unique UNIQUE (username);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
