MIGRATIONS_DIR = "./repo/db/migrations" 
DB_URL = "postgres://postgres:pg@localhost:5432/merema?sslmode=disable"
SCHEMA_DIR = "./repo/db/schema.sql"

run:
	export DB_URL=${DB_URL} && \
	go run .

setup:
	go get -u github.com/swaggo/swag
	go mod tidy
	docker compose up -d

create-migration: 
	dbmate -d ${MIGRATIONS_DIR} new ${name}

migrate-up:
	dbmate -d ${MIGRATIONS_DIR} -u ${DB_URL} -s ${SCHEMA_DIR} up

migrate-down:
	dbmate -d ${MIGRATIONS_DIR} -u ${DB_URL} -s ${SCHEMA_DIR} down

