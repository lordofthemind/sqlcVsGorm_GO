# Environment variables for SQLC database
SQLC_PG_CONTAINER_NAME ?= SqlcVsGorm_PG
SQLC_PG_IMAGE_TAG ?= latest
SQLC_PG_DB_NAME ?= SqlcVsGorm_SQLC
SQLC_PG_DB_USERNAME ?= postgres
SQLC_PG_DB_PASSWORD ?= postgresSqlcVsGormSecret
SQLC_PG_PORT ?= 5434
SQLC_PG_INTERNAL_PORT ?= 5432

# Environment variables for GORM database
GORM_PG_DB_NAME ?= SqlcVsGorm_GORM

# Colors for help command
CYAN := \033[36m
RESET := \033[0m

# Docker PostgreSQL commands
crtpg: ## Create and start the PostgreSQL container
	@echo "Creating and starting PostgreSQL container..."
	docker run --name $(SQLC_PG_CONTAINER_NAME) -p $(SQLC_PG_PORT):$(SQLC_PG_INTERNAL_PORT) -e POSTGRES_PASSWORD=$(SQLC_PG_DB_PASSWORD) -d postgres:$(SQLC_PG_IMAGE_TAG)

strpg: ## Start the PostgreSQL container
	@echo "Starting PostgreSQL container..."
	docker start $(SQLC_PG_CONTAINER_NAME)

stppg: ## Stop the PostgreSQL container
	@echo "Stopping PostgreSQL container..."
	docker stop $(SQLC_PG_CONTAINER_NAME)

rmvpg: ## Remove the PostgreSQL container
	@echo "Removing PostgreSQL container..."
	docker rm $(SQLC_PG_CONTAINER_NAME)

# PostgreSQL database commands for SQLC
crtpgdb_sqlc: strpg ## Create PostgreSQL database for SQLC
	@echo "Creating PostgreSQL database for SQLC..."
	docker exec -it $(SQLC_PG_CONTAINER_NAME) createdb -U $(SQLC_PG_DB_USERNAME) $(SQLC_PG_DB_NAME)

drppgdb_sqlc: strpg ## Drop PostgreSQL database for SQLC
	@echo "Dropping PostgreSQL database for SQLC..."
	docker exec -it $(SQLC_PG_CONTAINER_NAME) dropdb -U $(SQLC_PG_DB_USERNAME) $(SQLC_PG_DB_NAME)

# PostgreSQL database commands for GORM
crtpgdb_gorm: strpg ## Create PostgreSQL database for GORM
	@echo "Creating PostgreSQL database for GORM..."
	docker exec -it $(SQLC_PG_CONTAINER_NAME) createdb -U $(SQLC_PG_DB_USERNAME) $(GORM_PG_DB_NAME)

drppgdb_gorm: strpg ## Drop PostgreSQL database for GORM
	@echo "Dropping PostgreSQL database for GORM..."
	docker exec -it $(SQLC_PG_CONTAINER_NAME) dropdb -U $(SQLC_PG_DB_USERNAME) $(GORM_PG_DB_NAME)

# Help command
help: ## Show this help message
	@echo "Available commands:"
	@awk 'BEGIN {FS = ":.*##"; printf "\n\033[1m%-20s\033[0m %s\n\n", "Command", "Description"} /^[a-zA-Z_-]+:.*?##/ { printf "\033[36m%-20s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

.PHONY: crtpg strpg stppg rmvpg crtpgdb_sqlc drppgdb_sqlc crtpgdb_gorm drppgdb_gorm help
