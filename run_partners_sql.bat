@echo off
echo Running Partners Data SQL...
echo.

REM Check if MySQL is available
mysql --version >nul 2>&1
if %errorlevel% neq 0 (
    echo MySQL command line client not found!
    echo Please install MySQL or use MySQL Workbench/phpMyAdmin instead.
    echo.
    echo You can also manually run the partners_data.sql file in your MySQL tool.
    pause
    exit /b 1
)

echo Please enter your MySQL credentials:
echo.

REM Get database connection details
set /p DB_HOST="Database host (default: localhost): "
if "%DB_HOST%"=="" set DB_HOST=localhost

set /p DB_USER="Username: "
if "%DB_USER%"=="" (
    echo Username is required!
    pause
    exit /b 1
)

set /p DB_NAME="Database name: "
if "%DB_NAME%"=="" (
    echo Database name is required!
    pause
    exit /b 1
)

echo.
echo Running SQL file...
mysql -h %DB_HOST% -u %DB_USER% -p %DB_NAME% < partners_data.sql

if %errorlevel% equ 0 (
    echo.
    echo ✅ Partners data inserted successfully!
    echo.
    echo Next steps:
    echo 1. Restart your Go backend
    echo 2. Check the CMS partners page: http://localhost:5173/cms/partners
    echo 3. Check the public partnerships page: http://localhost:5173/about/partnerships
) else (
    echo.
    echo ❌ Error running SQL file. Please check your credentials and try again.
)

echo.
pause
