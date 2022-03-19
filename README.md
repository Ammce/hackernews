Create Migration: migrate create -ext sql -dir adapters/postgres/migrations create_users

Run migrations: migrate -path adapters/postgres/migrations -database "postgres://postgres:postgres@localhost:5433/news?sslmode=disable" up

Generate Types: go run github.com/99designs/gqlgen generate
