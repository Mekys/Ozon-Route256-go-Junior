# build docker image
build:
	docker-compose build

# run docker image
up-db:
	docker-compose up -d postgres

stop-db:
	docker-compose stop postgres

start-db:
	docker-compose start postgres

down-db:
	docker-compose down postgres



