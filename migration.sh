goose -dir ./migrations postgres "postgres://postgres:password@localhost:5433/postgres?sslmode=disable" status

goose -dir ./migrations postgres "postgres://postgres:password@localhost:5433/postgres?sslmode=disable" up