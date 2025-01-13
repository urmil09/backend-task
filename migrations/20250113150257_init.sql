-- +goose Up
-- +goose StatementBegin
SELECT 'select 1';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'select 2';
-- +goose StatementEnd
