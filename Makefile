dc: 
	docker-compose up;
swag gen: 
	swag init -o ./docs --parseInternal 
mock gen handlers:
	mockery --dir ./internal/handlers --all --output ./internal/handlers/mock --with-expecter
test:
	go  test -timeout 30s ./internal/repository ./internal/handlers/
lint:
	golangci-lint run --tests false ./