NETWORK_NAME := net101

.PHONY: init
init:
	@docker network create inspect $(NETWORK_NAME) >/dev/null 2>&1 || docker network create $(NETWORK_NAME)

.PHONY: run dev
run dev:
	docker compose -f docker-compose.dev.yaml up --build

migrate create:
	migrate create -ext .go -dir ./migrations -seq create_users_table