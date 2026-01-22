@echo off
echo Running Photo Gallery Migration...
echo.

mysql -u root -p gcx_cms < "database/migrations/2025_01_20_create_photo_galleries_table.sql"

if %errorlevel% equ 0 (
    echo.
    echo Migration completed successfully!
    echo Photo galleries and gallery photos tables created.
) else (
    echo.
    echo Migration failed!
    echo Please check the error messages above.
)

pause

