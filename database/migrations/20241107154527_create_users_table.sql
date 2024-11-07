-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
  id INT AUTO_INCREMENT NOT NULL,
  name VARCHAR(255),
  email VARCHAR(255),
  password VARCHAR(255),
  PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
