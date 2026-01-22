@echo off
echo Fixing Photo Gallery Counts...
echo.
echo This will update photo_count to match actual photos in each gallery.
echo.

mysql -u root -p gcx_cms < "database/migrations/fix_gallery_photo_counts.sql"

if %errorlevel% equ 0 (
    echo.
    echo Photo counts fixed successfully!
    echo All galleries now show correct photo counts.
) else (
    echo.
    echo Fix failed!
    echo Please check the error messages above.
)

pause

