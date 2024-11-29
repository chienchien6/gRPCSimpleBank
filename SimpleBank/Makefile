postgres:
	docker run --name pg -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=123456 -p 5432:5432
createdb:
	docker exec -it pg createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it pg dropdb --username=postgres simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://postgres:123456@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:123456@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server
