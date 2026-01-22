@echo off
echo Running Video Library Migration...
echo.

mysql -u root -p gcx_cms < "database/migrations/2025_01_20_create_video_library_table.sql"

if %errorlevel% equ 0 (
    echo.
    echo Migration completed successfully!
    echo Video libraries and library videos tables created.
) else (
    echo.
    echo Migration failed!
    echo Please check the error messages above.
)

pause
