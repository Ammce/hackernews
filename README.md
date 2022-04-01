### Hackernews Golang Graphql

## Purpose
- Purpose of this project is to try out how GraphQL works with Golang. 

## Scripts
- Start app `docker-compose up`
- Stop app `docker-compose down`
- Run Migrations `migrate -path adapters/postgres/migrations -database "postgres://postgres:postgres@localhost:5433/news?sslmode=disable" up`
- Generate GQL types `go run github.com/99designs/gqlgen generate`
- Create new migration `migrate create -ext sql -dir adapters/postgres/migrations create_users`

_Note:_ If you end up seeing this message when runnig migration: _error: Dirty database version 20220327002398. Fix and force version._, that means you did something wrong in migration. In order to resolve this, here are steps you need to do: 
- Find a migration that was run before that one. (This is easy, you just find previous timestamp migration)
- Run the fix: `migrate -path adapters/postgres/migrations -database "postgres://postgres:postgres@localhost:5433/news?sslmode=disable" force 20220328002314` (20220328002314 one is from the migration files, )

## Things left to do
_Please note this is just GraphQL test with Golang_
- Have better naming and code refactor
- Error handling
- Tests
- Subscriptions
- Generics
