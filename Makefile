# Load environment variables from config-local.env
ifneq (,$(wildcard config-local.env))
  include config-local.env
  export
endif

# Docker images
POSTGRES_IMAGE=postgres:15
PGADMIN_IMAGE=dpage/pgadmin4
LIQUIBASE_IMAGE=liquibase/liquibase

# Docker container names
POSTGRES_CONTAINER=my-postgres
PGADMIN_CONTAINER=pgadmin

# Run Liquibase update
liquibase-update:
	docker run --rm \
		-v $(PWD):/liquibase/changelog \
		-w /liquibase/changelog \
		$(LIQUIBASE_IMAGE) \
		--url=jdbc:postgresql://$(DB_HOST_LIQUIBASE):$(DB_PORT)/$(DB_NAME) \
		--username=$(DB_USER) \
		--password=$(DB_PASS) \
		--changeLogFile=migration/changelog.xml \
		update

# Run Liquibase status
liquibase-status:
	docker run --rm \
		-v $(PWD):/liquibase/changelog \
		-w /liquibase/changelog \
		$(LIQUIBASE_IMAGE) \
		--url=jdbc:postgresql://$(DB_HOST):$(DB_PORT)/$(DB_NAME) \
		--username=$(DB_USER) \
		--password=$(DB_PASS) \
		--changeLogFile=changelog.xml \
		status

run-postgres:
	docker run -d \
		--name $(POSTGRES_CONTAINER) \
		--env-file config-local.env \
		-e POSTGRES_USER=$(DB_USER) \
		-e POSTGRES_PASSWORD=$(DB_PASS) \
		-e POSTGRES_DB=$(DB_NAME) \
		-p 5432:5432 \
        -v pgdata:/var/lib/postgresql/data \
        postgres

# Start pgAdmin
run-pgadmin:
	docker run -d \
		--name $(PGADMIN_CONTAINER) \
		--env-file config-local.env \
		-e PGADMIN_DEFAULT_EMAIL=$(PGADMIN_EMAIL) \
		-e PGADMIN_DEFAULT_PASSWORD=$(PGADMIN_PASSWORD) \
		-p $(PGADMIN_PORT):80 \
		$(PGADMIN_IMAGE)

wait-for-postgres:
	@echo "Waiting for Postgres to be ready..."
	@sleep 10  # Adjust as needed

run-go:
	go run cmd/main.go

run:
	$(MAKE) liquibase-update
	$(MAKE) run-go

test:
	go test -v ./...

build:
	go build -o myapp cmd/main.go