APP_NAME=httpserver
ADDR=":8089"
VERSION="v0.0.1"

all: fmt vet test testrace

run: build-app
	@echo "run app..."
	@chmod +x ./$(APP_NAME)
	@./$(APP_NAME) -addr=$(ADDR) -version=$(VERSION)

build-image: all
	@echo "build image..."
	@docker build -t cncamp/$(APP_NAME) -f ./Dockerfile .

run-image: build-image
	@echo "run image..."
	@docker run -p 8089:8089 cncamp/$(APP_NAME) -addr=$(ADDR) -version=$(VERSION)

build-app: all
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
		build-app \
		build-image \
		run-image \
		run
