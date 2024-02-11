postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root myclass

dropdb:
	docker exec -it postgres16 dropdb myclass

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/myclass?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/myclass?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

mock:
	mockgen -build_flags=--mod=mod -package mockdb -destination db/mock/store.go myclass/db/sqlc

.PHONY: postgres createdb dropdb migrateup migratedown sqlc mock