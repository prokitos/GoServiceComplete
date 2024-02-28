-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS test (
 id SERIAL PRIMARY KEY,
 name VARCHAR(255) NOT NULL,
 surname VARCHAR(255) NOT NULL,
 patronymic VARCHAR(255) NOT NULL,
 age VARCHAR(255),
 sex VARCHAR(255),
 nationality VARCHAR(255)
);

INSERT INTO "test" ("id", "name", "surname", "patronymic","age","sex","nationality") VALUES (1, 'denis', 'denisov', 'denisovich','50','male','RU');
INSERT INTO "test" ("id", "name", "surname", "patronymic","age","sex","nationality") VALUES (2, 'Bob', 'Cannet', 'Johson','85','male','RU');
INSERT INTO "test" ("id", "name", "surname", "patronymic","age","sex","nationality") VALUES (3, 'denis', 'denisov', 'denisovich','50','male','RU');



-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS test;
-- +goose StatementEnd
