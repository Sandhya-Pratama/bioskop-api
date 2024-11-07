postgres:
	docker run --name postgres17 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres:latest

createdb:
	docker exec -it postgres17 createdb --username=root --owner=root bioskop-api

dropdb:
	docker exec -it postgres17 dropdb bioskop-api

createmigrations:
	migrate create -ext sql -dir db/migration -seq dbname

sqlc:
	sqlc generate

test: 
	go test -v -cover ./...

.PHONY: postgres createdb dropdb createmigrations sqlc test