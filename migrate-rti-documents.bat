@echo off
REM Migrate RTI documents table to database

echo ======================================
echo   RTI Documents Table Migration
echo ======================================
echo.
echo Creating rti_documents table...
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

mysql -h %DB_HOST% -P %DB_PORT% -u %DB_USER% -p%DB_PASS% %DB_NAME% < database\migrations\2025_01_20_create_rti_documents_table.sql

if %ERRORLEVEL% EQU 0 (
    echo.
    echo ======================================
    echo   Migration completed successfully!
    echo ======================================
    echo.
    echo The rti_documents table has been created with sample documents.
    echo.
    echo Sample documents included:
    echo   - GCX RTI Manual
    echo   - RTI Application Form
    echo   - RTI Guidelines
    echo   - RTI Act 2019
    echo   - Exemptions Guide
    echo.
    echo You can now:
    echo   1. Manage documents in CMS at /cms/rti (Documents tab)
    echo   2. Users can download from /rti (Resources tab)
    echo.
) else (
    echo.
    echo ======================================
    echo   Migration failed!
    echo ======================================
    echo.
)

pause
