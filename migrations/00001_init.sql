-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE users (
    id BIGINT,
    vk_id BIGINT
);

ALTER TABLE users 
ADD CONSTRAINT pk_users PRIMARY KEY (id),
ADD CONSTRAINT check_users__vk_id CHECK (vk_id > 0);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
