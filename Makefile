# Define the migrate command
MYSQL_URL = mysql://root:password@tcp(localhost:3306)/chat-app?charset=utf8mb4&parseTime=True&loc=Local

# Create the migrations folder if it doesn't exist
$(shell mkdir -p db/migrations)

# Define the migration targets
migrate-up: ## Apply all up migrations
	migrate -database "${MYSQL_URL}" -path db/migrations up

migrate-down: ## Revert all down migrations
	migrate -database "${MYSQL_URL}" -path db/migrations down

migrate-reset: ## Reset database to initial state
	migrate -database "${MYSQL_URL}" -path db/migrations drop

migrate-new: ## Create a new migration file
ifndef name
	$(error name is not set. Use 'make migrate-new name=create_user_table')
endif
	@echo "Creating new migration with name: $(name)"
	@cd db/migrations && migrate create -ext sql -seq $(name)

migrate-status: ## Show current migration status
	migrate -database "${MYSQL_URL}" -path db/migrations status

migrate-version: ## Show migrate version
	migrate -database "${MYSQL_URL}" -path db/migrations version

swag-generate: ## Generate Swagger Doc
	swag init --parseDependency --parseInternal --parseDepth 1 -g cmd/web/main.go

air-run: ## Run Air Live Reload
	air

deploy:
	docker-compose -p "chat-app-container" up -d

run-test:
	go test -cover ./...

help: ## Show this help message
	@echo "Available commands:"
	@awk 'BEGIN {FS = ":.*?## "}; /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: migrate-up migrate-down migrate-reset migrate-new migrate-status migrate-version migrate-help
