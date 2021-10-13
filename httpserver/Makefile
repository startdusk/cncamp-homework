APP_NAME=httpserver
ADDR=80
VERSION="v0.0.1"

TAG="v1.0"

all: fmt vet test testrace

run: build-app
	@echo "run app..."
	@chmod +x ./$(APP_NAME)
	@sudo ./$(APP_NAME) -addr=:$(ADDR) -version=$(VERSION)

build-image: all
	@echo "build image..."
	@docker build -t cncamp/$(APP_NAME) -f ./Dockerfile .

upload-image: build-image
	@echo "upload image $(APP_NAME)..."
	@docker tag cncamp/$(APP_NAME) cncamp/$(APP_NAME):$(TAG)
	@docker push cncamp/$(APP_NAME):$(TAG)

run-image: build-image
	@echo "run image..."
	@docker run -p $(ADDR):$(ADDR) cncamp/$(APP_NAME) -addr=:$(ADDR) -version=$(VERSION)

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
		upload-image \
		run
