# Makefile for IDOR in NFT Minting Lab

# Docker Compose file path
COMPOSE_FILE=./deployment/dev.docker-compose.yaml

# Default target
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  make dev       - Start the development environment using Docker Compose"
	@echo "  make stop      - Stop the running Docker Compose containers"
	@echo "  make restart   - Restart the environment"
	@echo "  make logs      - Show logs from all services"
	@echo "  make clean     - Remove containers and volumes (DANGER: will erase data)"

.PHONY: dev
dev:
	docker compose -f $(COMPOSE_FILE) up --build

.PHONY: stop
stop:
	docker compose -f $(COMPOSE_FILE) down

.PHONY: restart
restart:
	docker compose -f $(COMPOSE_FILE) down
	docker compose -f $(COMPOSE_FILE) up --build

.PHONY: logs
logs:
	docker compose -f $(COMPOSE_FILE) logs -f

.PHONY: clean
clean:
	docker compose -f $(COMPOSE_FILE) down -v
