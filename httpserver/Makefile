APP_NAME=httpserver
ADDR=":8089"
VERSION="v0.0.2"

all: fmt vet test testrace

run: build
	@echo "run app..."
	@chmod +x ./$(APP_NAME)
	@./$(APP_NAME) -addr=$(ADDR) -version=$(VERSION)

build: all
	@rm -rf $(APP_NAME) &>/dev/null
	@echo "build app..."
	@go build .

fmt:
	@echo "fmt code..."
	@go fmt ./...

vet:
	@echo "vet code..."
	@go vet ./...

test:
	@echo "run testing..."
	@go test ./...

testrace:
	@echo "run test race..."
	@go test -race -cpu 1,4 -timeout 7m ./...

.PHONY: \
		fmt \
		vet \
		test \
		testrace \
		build \
		run
