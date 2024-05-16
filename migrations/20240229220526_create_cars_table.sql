-- +goose Up
-- +goose StatementBegin
CREATE TABLE cars (
    id int PRIMARY KEY
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP table cars
-- +goose StatementEnd
