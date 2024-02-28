migrate-up:
	@migrate -path database/migrations -database "${DB_URI_PROD}" -verbose up
migrate-down:
	@migrate -path database/migrations -database "${DB_URI_PROD}" -verbose down
migrate-force:
	@read -p "Enter migration version: " version; \
		migrate -path database/migrations  -database "${DB_URI_PROD}" -verbose force $$version
create-migrate:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir database/migrations -seq $$name
run:
	@go run .
seed_user:
	@go run . seed_user
seed_trivia:
	@go run . seed_trivia
sql:
	@sqlc generate
build:
	@go build -o ./bin/trivia .

.PHONY: migrat-up migrate-down migrate-force create-migrate run sql
