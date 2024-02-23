COMPOSE_FILES=docker-compose.yml
COMPOSE_COMMAND=docker-compose

ifeq (, $(shell which $(COMPOSE_COMMAND)))
	COMPOSE_COMMAND=docker-compose
	ifeq (, $(shell which $(COMPOSE_COMMAND)))
		$(error "No docker compose in path, consider installing docker on your machine.")
	endif
endif

ifneq (,$(wildcard ./.env))
    include .env
    export
endif

env:
	@[ -e ./.env ] || cp -v ./.env.example ./.env

build:
	go build -v -ldflags "-s" -o build/ ./cmd/...
server:
	go run main.go server
lint:
	golangci-lint run
up:
	$(COMPOSE_COMMAND) -f $(COMPOSE_FILES) up -d
log:
	$(COMPOSE_COMMAND) -f $(COMPOSE_FILES) logs -f $(RUN_ARGS)
down:
	$(COMPOSE_COMMAND) -f $(COMPOSE_FILES) down --remove-orphans

purge:
	$(COMPOSE_COMMAND) -f $(COMPOSE_FILES) down --remove-orphans --volumes
status:
	$(COMPOSE_COMMAND) -f $(COMPOSE_FILES) ps
mysql-shell:
	$(COMPOSE_COMMAND) -f $(COMPOSE_FILES) exec -u 0 mysql mysql -hmysql -u$(MYSQL_USER) -D$(MYSQL_DATABASE) -p$(MYSQL_PASSWORD)
