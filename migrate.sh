#!/bin/bash
echo "Running database migration..."
go run cmd/migrate/main.go
