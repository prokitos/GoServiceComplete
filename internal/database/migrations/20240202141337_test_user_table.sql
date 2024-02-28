-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS testUser (
 id SERIAL PRIMARY KEY,
 name VARCHAR(255) NOT NULL,
 surname VARCHAR(255) NOT NULL,
 patronymic VARCHAR(255) NOT NULL,
 age VARCHAR(255),
 sex VARCHAR(255),
 nationality VARCHAR(255)
);

UPSERT INTO testUser (id, name, surname, patronymic, age, sex, nationality) VALUES
(1, 'Bob', 'Cannet', 'Johson', '85', 'male', 'RU'),
(3, 'denis', 'denisov', 'denisov', '50', 'male', 'RU');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS testUser;
-- +goose StatementEnd
