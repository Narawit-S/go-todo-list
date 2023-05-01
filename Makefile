postgres:
	docker run --name go-todo-list -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14.7

createdb:
	docker exec -it go-todo-list createdb --username=root --owner=root todo

dropdb:
	docker exec -it go-todo-list dropdb todo

migrateup:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5433/todo?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5433/todo?sslmode=disable" -verbose down

sqlc:
	sqlc generate
