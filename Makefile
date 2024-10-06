build:
	@echo "Building..."
	@go build -o dist/redis-clone cmd/main.go
	@echo "Build completed"
