include .env
export

postgresinit:
	docker run --name postgres-chat -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

psql:
	docker exec -it postgres-chat psql

createdb:
	docker exec -it postgres-chat createdb --username=root --owner=root postgres-chat-db

dropdb:
	docker exec -it postgres dropdb postgres-chat-db

migrateup:
	migrate -path db/migrations/ -database "$(DB_SOURCE)" --verbose up

migratedown:
	migrate -path db/migrations/ -database "$(DB_SOURCE)" --verbose down

.PHONY: postgresinit psql  createdb dropdb migrateup migratedown