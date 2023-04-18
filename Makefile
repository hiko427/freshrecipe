postgres:
	docker run --name myrecipe -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15.2-alpine
createdb:
	docker exec -it myrecipe createdb --username=root --owner=root myrecipe

dropdb:
	docker exec -it myrecipe dropdb myrecipe
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/myrecipe?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/myrecipe?sslmode=disable" -verbose down

sqlc:
	sqlc generate

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc	server	