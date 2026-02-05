#!/bin/bash
# Check MySQL databases and tables (run on VPS from gcx-backend folder)
# Usage: ./check-mysql-tables.sh   or   bash check-mysql-tables.sh

set -e

# Try to load from .env if present
if [ -f .env ]; then
  export $(grep -E '^DB_|^MYSQL_' .env | xargs 2>/dev/null) || true
fi

DB_NAME="${DB_NAME:-gcx_website}"
DB_USER="${DB_USER:-root}"
DB_PASS="${DB_PASSWORD:-gcxadmin123}"
DB_HOST="${DB_HOST:-127.0.0.1}"

echo "=========================================="
echo "MySQL connection: $DB_USER@$DB_HOST"
echo "Database: $DB_NAME"
echo "=========================================="
echo ""

echo "--- Databases ---"
mysql -h "$DB_HOST" -u "$DB_USER" -p"$DB_PASS" -e "SHOW DATABASES;" 2>/dev/null || mysql -h "$DB_HOST" -u "$DB_USER" -p"$DB_PASS" -e "SHOW DATABASES;"

echo ""
echo "--- Tables in $DB_NAME ---"
mysql -h "$DB_HOST" -u "$DB_USER" -p"$DB_PASS" "$DB_NAME" -e "SHOW TABLES;" 2>/dev/null || mysql -h "$DB_HOST" -u "$DB_USER" -p"$DB_PASS" "$DB_NAME" -e "SHOW TABLES;"

echo ""
echo "--- Key tables (row counts) ---"
for t in news_items events blog_posts pages settings users; do
  count=$(mysql -h "$DB_HOST" -u "$DB_USER" -p"$DB_PASS" "$DB_NAME" -N -e "SELECT COUNT(*) FROM \`$t\`;" 2>/dev/null) || count="(table missing or error)"
  echo "  $t: $count"
done

echo ""
echo "Done."
