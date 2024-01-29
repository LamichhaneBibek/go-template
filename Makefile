postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=lamichhane -d postgres:15-alpine
createdb:
	docker exec -it postgres15 createdb --username=postgres --owner=postgres template

dropdb:
	docker exec -it postgres15 dropdb --username=postgres  template

server:
	go run cmd/main.go

