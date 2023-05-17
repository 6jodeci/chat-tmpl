postgresinit:
	docker run --name postgres-chat -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it postgres-chat createdb --username=root --owner=root postgres-chat-db

dropdb:
	docker exec -it postgres dropdb postgres-chat-db

.PHONY: postgresinit createdb