run:
	@go run cmd/api/main.go

compose/start:
	@docker-compose up -d

compose/down:
	@docker-compose down