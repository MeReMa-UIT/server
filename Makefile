MIGRATIONS_DIR = "./repo/migrations" 
DB_URL = "postgres://postgres:pg@localhost:5432/merema?sslmode=disable"
SCHEMA_DIR = "./repo/schema.sql"

create-migration: 
	dbmate -d ${MIGRATIONS_DIR} new ${name}

migrate-up:
	dbmate -d ${MIGRATIONS_DIR} -u ${DB_URL} -s ${SCHEMA_DIR} up

migrate-down:
	dbmate -d ${MIGRATIONS_DIR} -u ${DB_URL} -s ${SCHEMA_DIR} down