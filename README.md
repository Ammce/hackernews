Create Migration: migrate create -ext sql -dir adapters/postgres/migrations create_users

Run migrations: migrate -path adapters/postgres/migrations -database "postgres://postgres:postgres@localhost:5433/news?sslmode=disable" up

Generate Types: go run github.com/99designs/gqlgen generate

If I have error about dirty migrations: "error: Dirty database version 16. Fix and force version.", should force migrate version before it.

migrate -path adapters/postgres/migrations -database "postgres://postgres:postgres@localhost:5433/news?sslmode=disable" force 20220328002314 (this one is from the migration files)
