build:
	@go build -o bin/ecom cmd/main.go

test:
	@go test -v ./...

run:
	@go run cmd/main.go


	
# Create a new database migration file with SQL extension
# Usage: make migration <migration_name>
# Example: make migration create_users_table
# This will create two files in cmd/migrate/migrations/:
# - YYYYMMDDHHMMSS_<migration_name>.up.sql (for applying changes)
# - YYYYMMDDHHMMSS_<migration_name>.down.sql (for reverting changes)
# The @ symbol suppresses command echoing for cleaner output
migration:
	@C:\Users\nutan\go\bin\migrate.exe create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))
migrate-up:
	@go run cmd/migrate/main.go up	
migrate-down:
	@go run cmd/migrate/main.go down

# Check current migration version and status
migrate-status:
	@echo "Checking migration status..."
	@C:\Users\nutan\go\bin\migrate.exe -database "mysql://root:@tcp(127.0.0.1:3306)/ecom" -path cmd/migrate/migrations version

# Fix dirty migration by forcing to a clean state
# Usage: make migrate-fix-dirty VERSION=20241017123456
migrate-fix-dirty:
	@echo "Fixing dirty migration..."
	@if "$(VERSION)" == "" (echo "Error: Please provide VERSION. Usage: make migrate-fix-dirty VERSION=20241017123456" && exit 1)
	@C:\Users\nutan\go\bin\migrate.exe -database "mysql://root:@tcp(127.0.0.1:3306)/ecom" -path cmd/migrate/migrations force $(VERSION)
	@echo "Migration forced to version $(VERSION). Now run 'make migrate-up' to continue."

# Auto-fix dirty migration (gets current dirty version and forces it clean)
migrate-fix-auto:
	@echo "Auto-fixing dirty migration..."
	@for /f "tokens=1" %%i in ('C:\Users\nutan\go\bin\migrate.exe -database "mysql://root:@tcp(127.0.0.1:3306)/ecom" -path cmd/migrate/migrations version 2^>nul ^| findstr /r "[0-9].*dirty"') do C:\Users\nutan\go\bin\migrate.exe -database "mysql://root:@tcp(127.0.0.1:3306)/ecom" -path cmd/migrate/migrations force %%i
	@echo "Dirty migration fixed. Now you can run 'make migrate-up'."

# Go to specific migration version
# Usage: make migrate-goto VERSION=20241017123456
migrate-goto:
	@if "$(VERSION)" == "" (echo "Error: Please provide VERSION. Usage: make migrate-goto VERSION=20241017123456" && exit 1)
	@C:\Users\nutan\go\bin\migrate.exe -database "mysql://root:@tcp(127.0.0.1:3306)/ecom" -path cmd/migrate/migrations goto $(VERSION)

# Apply one migration at a time (useful for debugging)
migrate-up-one:
	@C:\Users\nutan\go\bin\migrate.exe -database "mysql://root:@tcp(127.0.0.1:3306)/ecom" -path cmd/migrate/migrations up 1

# Rollback one migration at a time
migrate-down-one:
	@C:\Users\nutan\go\bin\migrate.exe -database "mysql://root:@tcp(127.0.0.1:3306)/ecom" -path cmd/migrate/migrations down 1

# Reset all migrations (DANGEROUS - drops all tables)
migrate-reset:
	@echo "WARNING: This will drop all tables and reset migrations!"
	@set /p confirm="Are you sure? Type 'yes' to continue: "
	@if "$(confirm)" == "yes" (C:\Users\nutan\go\bin\migrate.exe -database "mysql://root:@tcp(127.0.0.1:3306)/ecom" -path cmd/migrate/migrations drop) else (echo "Operation cancelled.")

# Show database tables and structure
db-status:
	@go run cmd/check/main.go

# Help for migration commands
migrate-help:
	@echo "Available migration commands:"
	@echo "  make migration NAME          - Create new migration files"
	@echo "  make migrate-up             - Apply all pending migrations"
	@echo "  make migrate-down           - Rollback all migrations"
	@echo "  make migrate-status         - Check current migration status"
	@echo "  make migrate-fix-dirty VERSION=123 - Fix dirty migration with specific version"
	@echo "  make migrate-fix-auto       - Auto-fix dirty migration"
	@echo "  make migrate-goto VERSION=123 - Go to specific migration version"
	@echo "  make migrate-up-one         - Apply only one migration"
	@echo "  make migrate-down-one       - Rollback only one migration"
	@echo "  make migrate-reset          - Reset all migrations (DANGEROUS)"
	@echo "  make db-status              - Show database tables"