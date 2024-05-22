postgres:
	docker run --name swd-postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -e POSTGRES_DB=dental_clinic -d postgres:15-alpine

migrate-up:
	migrate -path db/migrations -database "postgres://root:secret@localhost:5432/dental_clinic?sslmode=disable" -verbose up

migrate-down:
	migrate -path db/migrations -database "postgres://root:secret@localhost:5432/dental_clinic?sslmode=disable" -verbose down

sqlc:
	sqlc generate

server:
	go run main.go
