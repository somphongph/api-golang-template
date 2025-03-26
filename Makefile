dev:
	@echo "Running..."
	@air
	
## Test the application
test:
	@echo "Testing..."
	@go test ./... -v


## Integrations Tests for the application
test-in:
	@echo "Running integration tests..."
	@go test ./internal/database -v

## Test cover
test-cover:
	@echo "Testing with coverage..."
	@go test -coverpayment-gateways=coverage.out ./...
	@go tool cover -html=coverage.out

## Clean the binary
clean:
	@echo "Cleaning..."
	@ -f main

uplib:
	@echo "Upgrading libraries..."
	@go get -u github.com/somphongph/lib-golang-packages@latest

uplibdev:
	@echo "Upgrading libraries..."
	@go get -u github.com/somphongph/lib-golang-packages@develop