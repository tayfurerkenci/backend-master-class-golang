postgres:
	docker run --name postgres15.4 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=tayfur -d postgres:15.4-alpine
createdb:
	docker exec -it postgres15.4 createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it postgres15.4 dropdb simple_bank
migrateup: 
	migrate -path db/migration -database "postgresql://root:tayfur@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:tayfur@localhost:5432/simple_bank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
.PHONY: postgres createdb dropdb migrateup migratedown sqlc
