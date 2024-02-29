SERVICE_NAME=cars

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
swagger/gen:
	@rm -rf ./cmd/http
	@mkdir -p ./cmd/http

	@rm -rf ./internal/port/http/generated
	@mkdir -p ./internal/port/http/generated

	@go-swagger generate server \
		-f ./docs/openapi.yaml \
		-t ./internal/port/http/generated \
		-C ./swagger-templates/default-server.yml \
		--template-dir ./swagger-templates/templates \
		--name ${SERVICE_NAME}
