build:
	docker-compose --env-file ./app/development.env build

buildup:
	docker-compose --env-file ./app/development.env up --build

up:
	docker-compose --env-file ./app/development.env up

rm:
	docker-compose --env-file ./app/development.env rm
