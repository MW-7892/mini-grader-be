#!make

include .env

migrate_up:
	goose -v 									 \
		-dir database/migrations \
		mysql 									 \
		"${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" \
		up

migrate_down:
	goose -v  								 \
		-dir database/migrations \
		mysql 									 \
		"${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" \
		down
	
