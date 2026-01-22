@echo off
REM Migrate RTI requests table to database

echo ======================================
echo   RTI Requests Table Migration
echo ======================================
echo.
echo Creating rti_requests table...
echo.

REM Load environment variables from .env if it exists
if exist .env (
    for /F "tokens=*" %%i in ('type .env ^| findstr /v "^#"') do set %%i
)

REM Set default values if not provided
if "%DB_HOST%"=="" set DB_HOST=localhost
if "%DB_PORT%"=="" set DB_PORT=3306
if "%DB_NAME%"=="" set DB_NAME=gcx_cms
if "%DB_USER%"=="" set DB_USER=root
if "%DB_PASS%"=="" set DB_PASS=

mysql -h %DB_HOST% -P %DB_PORT% -u %DB_USER% -p%DB_PASS% %DB_NAME% < database\migrations\2025_01_20_create_rti_requests_table.sql

if %ERRORLEVEL% EQU 0 (
    echo.
    echo ======================================
    echo   Migration completed successfully!
    echo ======================================
    echo.
    echo The rti_requests table has been created with sample data.
    echo.
    echo You can now:
    echo   1. Access RTI requests in CMS at /cms/rti
    echo   2. Users can submit requests at /rti
    echo.
) else (
    echo.
    echo ======================================
    echo   Migration failed!
    echo ======================================
    echo.
)

pause
