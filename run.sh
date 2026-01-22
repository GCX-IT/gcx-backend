#!/bin/bash

# GCX CMS Go Backend Startup Script

echo "🚀 Starting GCX CMS Go Backend..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go 1.21 or higher."
    exit 1
fi

# Check if .env file exists, if not copy from example
if [ ! -f .env ]; then
    echo "📄 Creating .env file from example..."
    cp env.example .env
    echo "✅ .env file created. You can edit it to customize your settings."
fi

# Download dependencies
echo "📦 Downloading Go dependencies..."
go mod tidy

# Create uploads directory
echo "📁 Creating uploads directory..."
mkdir -p uploads/images

# Run the application
echo "🏃 Starting the server..."
echo "📡 Server will be available at: http://localhost:8080"
echo "🔐 Default admin login: admin@gcx.com / admin123"
echo ""
echo "🔄 To switch your frontend to use this backend, change BACKEND_TYPE to 'go' in frontend/src/plugins/axios.ts"
echo ""

go run main.go
