APP_NAME=Bookmark_Manager
BIN_DIR=gobin
BUILD_DIR=./bin
GO_FILES=$(shell find . -name '*.go' -not -path "./vendor/**")

run:
	@echo "Running the Server"
	@go run cmd/app/main.go

test:
	@echo "Running the Tests"
	@go test -v ./...

deps:
	@echo "Installing dependencies"
	@go mod tidy

build:
	@echo "Building the server"
	@mkdir -p ${BUILD_DIR}
	@go build -o $(BUILD_DIR)/$(BIN_DIR) cmd/app/main.go
	@echo "Build complete"

clean:
	@echo "Cleaning up"
	@rm -rf ${BUILD_DIR}
	@echo "Cleanup complete"

migrate-up:
	@echo "Migrating up"
	@bash -c 'source .env && goose -dir sql/migrations $$GOOSE_DRIVER $$GOOSE_DBSTRING up'


migrate-down:
	@echo "Migrating down"
	@bash -c 'source .env && goose -dir sql/migrations $$GOOSE_DRIVER $$GOOSE_DBSTRING down'