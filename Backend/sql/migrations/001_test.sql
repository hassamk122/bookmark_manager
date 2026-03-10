-- +goose Up
CREATE TABLE IF NOT EXISTS test_table(
    test_table_id SERIAL PRIMARY KEY
);

-- +goose Down
DROP TABLE test_table;