run: 
	@go mod vendor
	@echo "running service" 
	@go run internal/cmd/http/app.go 