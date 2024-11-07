postgres:
	docker run --name postgres17 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres:latest

createdb:
	docker exec -it postgres17 createdb --username=root --owner=root bioskop-api

dropdb:
	docker exec -it postgres17 dropdb bioskop-api

createmigrations:
	goose -dir db/migrations create create_user_table sql

migrateup:
	goose -dir db/migrations postgres "user=root dbname=bioskop-api sslmode=disable password=mysecretpassword port=5433" up

migratedown:
	goose -dir db/migrations postgres "user=root dbname=bioskop-api sslmode=disable password=mysecretpassword port=5433" down

.PHONY: postgres createdb dropdb createmigrations migrateup migratedown