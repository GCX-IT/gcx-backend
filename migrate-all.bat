@echo off
echo Running all migrations...
cd go-backend

echo Migrating board members...
go run cmd/migrate-board/main.go

echo Migrating brokers...
go run cmd/migrate-brokers/main.go

echo Migrating careers...
go run cmd/migrate-careers/main.go

echo Migrating commodities...
go run cmd/migrate-commodities/main.go

echo Migrating publications...
go run cmd/migrate-publications/main.go

echo Migrating social media...
go run cmd/migrate-social-media/main.go

echo Migrating team...
go run cmd/migrate-team/main.go

echo Migrating traders...
go run cmd/migrate-traders/main.go

echo All migrations completed!
pause
