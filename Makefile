.PHONY: db-up app-up down build migrate-up migrate-down service-start

db-up:
	docker compose up -d db
	@echo "DB started"

app-up:
	docker compose up -d app
	@echo "app started"

down:
	docker compose down

build:
	docker build -t qa-service-app -f Dockerfile .
	docker build -t qa-service-migrations -f Dockerfile.goose .

migrate-up:
	docker compose run --rm migrations goose up
	@echo "migrations rolled successfully"


migrate-down:
	docker compose run --rm migrations goose down
	@echo "migrations rolled-back successfully"

migrate-status:
	docker compose run --rm migrations goose status
	@echo "migrations status given successfully"

service-start: build db-up migrate-up app-up
	@echo "qa_service started on /localhost:8080"