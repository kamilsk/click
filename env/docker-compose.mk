COMPOSE ?= docker-compose -f env/docker-compose.base.yml -f env/docker-compose.dev.yml -p click

.PHONY: env
env:
	cp -n env/.env{.example,} || true # for containers
	cp -n env/.env .env       || true # for docker compose file, https://docs.docker.com/compose/env-file/

.PHONY: rm-env
rm-env:
	find . -name .env | xargs rm -f || true


.PHONY: config
config:
	$(COMPOSE) config

.PHONY: up
up: env
	$(COMPOSE) up -d
	$(COMPOSE) rm -f

.PHONY: fresh-up
fresh-up: env
	$(COMPOSE) up --build --force-recreate -d
	$(COMPOSE) rm -f

.PHONY: down
down: env
	$(COMPOSE) down

.PHONY: clean-down
clean-down: env
	$(COMPOSE) down --volumes --rmi local

.PHONY: clear
clear: env
	$(COMPOSE) rm -f

.PHONY: status
status: env
	$(COMPOSE) ps


.PHONY: up-db
up-db:
	$(COMPOSE) up -d db

.PHONY: start-db
start-db: env
	$(COMPOSE) start db

.PHONY: stop-db
stop-db: env
	$(COMPOSE) stop db

.PHONY: log-db
log-db: env
	$(COMPOSE) logs -f db

.PHONY: psql
psql: env
	$(COMPOSE) exec db /bin/sh -c 'su - postgres -c psql'

.PHONY: backup
backup: env
	$(COMPOSE) exec db \
	  /bin/sh -c 'su - postgres -c "pg_dump --clean $${POSTGRES_DB}"' > ./env/backup.sql

.PHONY: restore
restore:
	cat ./env/backup.sql | docker exec -i $$(make status | tail +3 | awk '{print $$1}' | grep _db_ | head -1) \
	  /bin/sh -c 'cat > /tmp/backup.sql && su - postgres -c "psql $${POSTGRES_DB} < /tmp/backup.sql"'


.PHONY: up-migration
up-migration:
	$(COMPOSE) up --build -d migration

.PHONY: start-migration
start-migration: env
	$(COMPOSE) start migration

.PHONY: log-migration
log-migration: env
	$(COMPOSE) logs -f migration


.PHONY: up-service
up-service:
	$(COMPOSE) up --build -d service

.PHONY: start-service
start-service: env
	$(COMPOSE) start service

.PHONY: stop-service
stop-service: env
	$(COMPOSE) stop service

.PHONY: log-service
log-service: env
	$(COMPOSE) logs -f service

.PHONY: demo
demo: env
	$(COMPOSE) exec service click migrate up --with-demo


.PHONY: up-server
up-server:
	$(COMPOSE) up -d server

.PHONY: start-server
start-server: env
	$(COMPOSE) start server

.PHONY: stop-server
stop-server: env
	$(COMPOSE) stop server

.PHONY: log-server
log-server: env
	$(COMPOSE) logs -f server
