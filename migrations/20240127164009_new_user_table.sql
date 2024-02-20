-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
 id SERIAL PRIMARY KEY,
 name VARCHAR(255) NOT NULL,
 surname VARCHAR(255) NOT NULL,
 patronymic VARCHAR(255) NOT NULL,
 age VARCHAR(255),
 sex VARCHAR(255),
 nationality VARCHAR(255)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
