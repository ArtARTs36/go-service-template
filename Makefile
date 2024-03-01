SERVICE_NAME=cars
API_PATH 		   = api/grpc/cars
PROTO_API_DIR 	   = api/grpc/cars
PROTO_OUT_DIR 	   = pkg/cars-grpc-api
PROTO_API_OUT_DIR  = ${PROTO_OUT_DIR}

# usage as: make migration/create NAME=create_cars_table
migration/create:
	docker run --rm \
		-v "${PWD}/migrations:/migrations" \
		artsafin/goose-migrations \
		create ${NAME} sql

# usage as: make migration/add-table NAME=cars
migration/add-table:
	docker run --rm \
		-v "${PWD}/migrations:/migrations" \
		artsafin/goose-migrations \
		create create_${NAME}_table sql

# usage as: make migrate
migration/migrate:
	docker-compose run migrations

# generate swagger structures
gen/swagger:
	@rm -rf ./cmd/http
	@mkdir -p ./cmd/http

	@rm -rf ./internal/port/http/generated
	@mkdir -p ./internal/port/http/generated

	@go-swagger generate server \
		-f ./api/openapi/openapi.yaml \
		-t ./internal/port/http/generated \
		-C ./swagger-templates/default-server.yml \
		--template-dir ./swagger-templates/templates \
		--name ${SERVICE_NAME}

gen/proto:
	mkdir -p ${PROTO_OUT_DIR}
	protoc \
		-I ${API_PATH}/v1 \
		--include_imports \
		--go_out=$(PROTO_OUT_DIR) --go_opt=paths=source_relative \
        --go-grpc_out=$(PROTO_OUT_DIR) --go-grpc_opt=paths=source_relative \
        --descriptor_set_out=$(PROTO_OUT_DIR)/api.pb \
		--go-client-constructor_out=$(PROTO_OUT_DIR) \
            --go-client-constructor_opt=generate-mock-client=true \
            --go-client-constructor_opt=embed-client=true \
        ./${PROTO_API_DIR}/v1/*.proto

gen/go:
	go generate ./...

gen: gen/go gen/proto gen/swagger
