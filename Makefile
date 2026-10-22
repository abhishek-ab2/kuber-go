-include .env
export

DOCKER_COMPOSE_DEPS=./infra/docker-compose-deps.yml

run-server:
	go run ./cmd/api

build-deps:
	docker compose -f $(DOCKER_COMPOSE_DEPS) build

start-deps:
	docker compose -f $(DOCKER_COMPOSE_DEPS) up -d

deps-down:
	docker compose -f $(DOCKER_COMPOSE_DEPS) down

deps-purge:
		docker compose -f $(DOCKER_COMPOSE_DEPS) down -v

build-and-start-deps:
	docker compose -f $(DOCKER_COMPOSE_DEPS) up -d --build
