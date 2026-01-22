@echo off
echo 🚀 Starting GCX CMS Go Backend...

REM Check if Go is installed
go version >nul 2>&1
if %errorlevel% neq 0 (
    echo ❌ Go is not installed. Please install Go 1.21 or higher.
    pause
    exit /b 1
)

REM Check if .env file exists, if not copy from example
if not exist .env (
    echo 📄 Creating .env file from example...
    copy env.example .env
    echo ✅ .env file created with MySQL configuration.
    echo ⚠️  Please edit .env file to add your MySQL password if needed.
    echo 📖 Check setup-mysql.md for detailed MySQL configuration instructions.
    pause
)

REM Download dependencies
echo 📦 Downloading Go dependencies...
go mod tidy

REM Create uploads directory
echo 📁 Creating uploads directory...
if not exist uploads\images mkdir uploads\images

REM Run the application
echo 🏃 Starting the server...
echo 📡 Server will be available at: http://localhost:8080
echo 🔐 Default admin login: admin@gcx.com / admin123
echo.
echo 🔄 To switch your frontend to use this backend, change BACKEND_TYPE to 'go' in frontend/src/plugins/axios.ts
echo.

go run main.go

pause
