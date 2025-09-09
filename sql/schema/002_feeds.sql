-- +goose Up
CREATE TABLE feeds(
  id  uuid PRIMARY KEY,
  name  TEXT NOT NULL,
  url TEXT UNIQUE NOT NULL,
  user_id uuid REFERENCES users(id) ON DELETE CASCADE,
  created_at  TIMESTAMP NOT NULL,
  update_at   TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE feeds;
