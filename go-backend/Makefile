postgres:
	docker run --name postgres15.4 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=tayfur -d postgres:15.4-alpine
createdb:
	docker exec -it postgres15.4 createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it postgres15.4 dropdb simple_bank
migrateup: 
	migrate -path db/migration -database "postgresql://root:tayfur@localhost:5432/simple_bank?sslmode=disable" -verbose up
migrateup1: 
	migrate -path db/migration -database "postgresql://root:tayfur@localhost:5432/simple_bank?sslmode=disable" -verbose up 1
migratedown:
	migrate -path db/migration -database "postgresql://root:tayfur@localhost:5432/simple_bank?sslmode=disable" -verbose down
migratedown1:
	migrate -path db/migration -database "postgresql://root:tayfur@localhost:5432/simple_bank?sslmode=disable" -verbose down 1
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/tayfurerkenci/simple-bank/go-backend/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc server mock
