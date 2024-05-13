run-consumer:
	@echo "Running consumer"
	@go run cmd/hotel-consumer/main.go

run-producer:
	@echo "Running producer"
	@go run cmd/hotel-provider/main.go