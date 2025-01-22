DB_URL=postgresql://postgres:123456@localhost:5432/simple_bank?sslmode=disable



postgres:
	docker run --name pg --network bank-network -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=123456 -p 5432:5432
createdb:
	docker exec -it pg createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it pg dropdb --username=postgres simple_bank

migrateup:
	migrate -path db/migration -database "${DB_URL}" -verbose up

migrateup1:
	migrate -path db/migration -database "${DB_URL}" -verbose up 1


migratedown:
	migrate -path db/migration -database "${DB_URL}" -verbose down

migratedown1:
	migrate -path db/migration -database "${DB_URL}" -verbose down 1


sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	 mockgen -package mockdb -destination db/mock/store.go github.com/GRPCgRPCBank/SimpleBank/db/sqlc Store

db_docs:
	 dbdocs build doc/db.dbml

db_scheme:
	dbml2sql --postgres -o doc/scheme.sql doc/db.dbml

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server migratedown1 migrateup1 mock db_docs db_scheme
