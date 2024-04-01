build:
	go build -o app .

test:
	go test ./...

run:
	go run main.go

postgres:
	docker run --name postgres-gorm2 -p 5432:5432 -e POSTGRES_PASSWORD=secret-gorm2 -d postgres

postgres-stop:
	docker stop postgres-gorm2

postgres-remove:
	docker rm postgres-gorm2

postgres-logs:
	docker logs postgres-gorm2

postgres-connect:
	docker exec -it postgres-gorm2 psql -U postgres

createdb:
	docker exec -it postgres-gorm2 createdb --username=postgres --owner=postgres person_type_gorm

create-test-db:
	docker exec -it postgres-gorm2 createdb --username=postgres --owner=postgres person_type_gorm_test

.PHONY: build test run postgres postgres-stop postgres-remove postgres-logs postgres-connect createdb