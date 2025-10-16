# PowerShell script to fix dirty migration state

$migratePath = "C:\Users\nutan\go\bin\migrate.exe"

Write-Host "Checking migration status..."

# Check current migration version
$checkVersionCmd = "`"$migratePath`" -database `"mysql://root:@tcp(127.0.0.1:3306)/ecom`" -path cmd/migrate/migrations version"
Write-Host "Running: $checkVersionCmd"
Invoke-Expression $checkVersionCmd

Write-Host "`nForcing migration to clean state..."

# Force the migration to version 20251015194357
$forceCmd = "`"$migratePath`" -database `"mysql://root:@tcp(127.0.0.1:3306)/ecom`" -path cmd/migrate/migrations force 20251015194357"
Write-Host "Running: $forceCmd"
Invoke-Expression $forceCmd

Write-Host "`nChecking migration status again..."
Invoke-Expression $checkVersionCmd

Write-Host "`nNow you can run 'make migrate-up' again."