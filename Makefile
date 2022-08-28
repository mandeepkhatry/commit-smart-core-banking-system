migrateup:
	migrate -path migrations/ -database "postgresql://postgres:admin@localhost:5432/commit_smart_cbs?sslmode=disable" -verbose up

migratedown:
	migrate -path migrations/ -database "postgresql://postgres:admin@localhost:5432/commit_smart_cbs?sslmode=disable" -verbose down

tidy:
	go mod tidy

run_server:
	go run cmd/main.go