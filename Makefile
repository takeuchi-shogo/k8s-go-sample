build:
	@echo "Building docker image."
	docker-compose --env-file ./app/development.env build
	@echo "Build docker image success."

buildup:
	@echo "Building docker image."
	docker-compose --env-file ./app/development.env up --build
	@echo "Build docker image success."

up:
	@echo "Running docker image."
	docker-compose --env-file ./app/development.env up
	@echo "Run docker image."

down:
	@echo "Stopping docker image."
	docker-compose --env-file ./app/development.env up
	@echo "Stop docker image."

rm:
	@echo "Removing docker image."
	docker-compose --env-file ./app/development.env rm
	@echo "Remove docker image success."
