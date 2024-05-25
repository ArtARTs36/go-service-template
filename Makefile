SERVICE_NAME=cars
API_PATH 		   = api/grpc/cars
PROTO_API_DIR 	   = api/grpc/cars
PROTO_OUT_DIR 	   = pkg/cars-grpc-api
PROTO_API_OUT_DIR  = ${PROTO_OUT_DIR}

.DEFAULT_GOAL := help

help: ## Show help
	@printf "\033[33m%s:\033[0m\n" 'Available commands'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_\/-]+:.*?## / {printf "  \033[32m%-18s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

tidy: ## Add missing and remove unused GO modules
	go mod tidy

migration/create: ## Create goose migration, usage as: make migration/create NAME=create_cars_table
	docker run --rm \
		-v "${PWD}/migrations:/migrations" \
		artsafin/goose-migrations \
		create ${NAME} sql

migration/add-table: ## Create migration for add table, usage as: make migration/add-table NAME=cars
	docker run --rm \
		-v "${PWD}/migrations:/migrations" \
		artsafin/goose-migrations \
		create create_${NAME}_table sql

migration/migrate: ## Migrate database schema
	docker-compose run migrations

gen/proto: ## Generate gRPC structures
	mkdir -p ${PROTO_OUT_DIR}
	protoc \
		-I ${API_PATH}/v1 \
		--include_imports \
		--go_out=$(PROTO_OUT_DIR) --go_opt=paths=source_relative \
        --go-grpc_out=$(PROTO_OUT_DIR) --go-grpc_opt=paths=source_relative \
        --descriptor_set_out=$(PROTO_OUT_DIR)/api.pb \
        ./${PROTO_API_DIR}/v1/*.proto

gen/go: ## Generate go/mock structures
	go generate ./...

gen: ## Generate go/mock, gRPC structures
	make gen/go
	make gen/proto

test: ## Run go tests
	go test ./...

lint: ## Run linter
	golangci-lint run --fix

up: ## Up services (foreground)
	docker-compose up -d postgres
	docker-compose run --rm migrations
	docker-compose run --rm cars-grpc

up-d: ## Up services (background)
	docker-compose up -d
