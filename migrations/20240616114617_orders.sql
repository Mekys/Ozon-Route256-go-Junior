-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS orders (
    orderId BIGSERIAL PRIMARY KEY,
    addresseeId BIGSERIAL NOT NULL,
    shelfLife TIMESTAMP NOT NULL,
    order_status INTEGER NOT NULL,
    status_updated_date TIMESTAMP,
    hash_code Text NOT NULL
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS order;
-- +goose StatementEnd
