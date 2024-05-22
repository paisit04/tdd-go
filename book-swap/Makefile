.PHONY: test mock start stop

build:
	docker-compose --env-file docker.env build --no-cache

start:
	docker-compose --env-file docker.env up

stop:
	docker-compose --env-file docker.env down

shell:
	docker-compose --env-file docker.env exec books bash

test:
	BOOKSWAP_DB_URL=postgres://root:root@localhost:5432/books?sslmode=disable go test -v ./...

mock:
	mockery --all