@echo off
echo Fixing Gallery Photo URLs...
echo.
echo This will update the sample photo URLs to use actual available images.
echo.

mysql -u root -p gcx_cms < "database/migrations/fix_gallery_photo_urls.sql"

if %errorlevel% equ 0 (
    echo.
    echo Photo URLs fixed successfully!
    echo Sample photos now use: trading.jpg, farmer.jpg, maize.jpg, crop.jpg
) else (
    echo.
    echo Fix failed!
    echo Please check the error messages above.
)

pause
