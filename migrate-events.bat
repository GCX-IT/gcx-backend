@echo off
REM Migrate events table to database
REM This script creates the events table and inserts sample data

echo ======================================
echo   GCX Events Table Migration
echo ======================================
echo.

REM Load environment variables from .env if it exists
if exist .env (
    echo Loading environment variables from .env...
    for /F "tokens=*" %%i in ('type .env ^| findstr /v "^#"') do set %%i
)

REM Set default values if not provided
if "%DB_HOST%"=="" set DB_HOST=localhost
if "%DB_PORT%"=="" set DB_PORT=3306
if "%DB_NAME%"=="" set DB_NAME=gcx_cms
if "%DB_USER%"=="" set DB_USER=root
if "%DB_PASS%"=="" set DB_PASS=

echo.
echo Database Configuration:
echo   Host: %DB_HOST%
echo   Port: %DB_PORT%
echo   Database: %DB_NAME%
echo   User: %DB_USER%
echo.

echo Running events table migration...
echo.

mysql -h %DB_HOST% -P %DB_PORT% -u %DB_USER% -p%DB_PASS% %DB_NAME% < database\migrations\2025_01_20_create_events_table.sql

if %ERRORLEVEL% EQU 0 (
    echo.
    echo ======================================
    echo   Migration completed successfully!
    echo ======================================
    echo.
    echo The events table has been created with sample data.
    echo You can now:
    echo   1. Start the backend server: go run main.go
    echo   2. Access the CMS at http://localhost:8080/cms
    echo   3. Manage events from the Events section
    echo.
) else (
    echo.
    echo ======================================
    echo   Migration failed!
    echo ======================================
    echo.
    echo Please check your database connection and try again.
    echo.
)

pause
