postgres:
	docker run --name myrecipe --network recipe-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15.2-alpine
createdb:
	docker exec -it myrecipe createdb --username=root --owner=root myrecipe

dropdb:
	docker exec -it myrecipe dropdb myrecipe
migrateup:
	migrate -path db/migration -database "postgresql://root:QBcjeHUHKn31qtirXKXe@myrecipe.cptum6fxzwpe.ap-northeast-1.rds.amazonaws.com:5432/recipe" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:QBcjeHUHKn31qtirXKXe@myrecipe.cptum6fxzwpe.ap-northeast-1.rds.amazonaws.com:5432/myrecipe" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc	server test	