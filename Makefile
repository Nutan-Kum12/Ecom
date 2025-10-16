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