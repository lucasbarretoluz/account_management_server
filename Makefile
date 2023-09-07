postgres:
	docker run --name postgres_a_m -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

createdb:
	docker exec -it postgres_a_m createdb --username=root --owner=root simple_account_management

dropdb:
	docker exec -it postgres_a_m dropdb simple_account_management

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_account_management?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_account_management?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server