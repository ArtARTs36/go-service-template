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
