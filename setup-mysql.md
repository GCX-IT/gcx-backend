# MySQL Configuration for GCX CMS

## 📋 Step 1: Create .env file

Copy `env.example` to `.env` and update it with your MySQL settings:

```bash
cp env.example .env
```

Edit `.env` file to look like this:

```env
# Database Configuration
DB_TYPE=mysql
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=gcx_website
DB_USER=root
DB_PASSWORD=your_mysql_password_here

# JWT Configuration
JWT_SECRET=gcx_super_secret_jwt_key_change_this_in_production
JWT_EXPIRES_IN=24h

# Server Configuration
PORT=8080
GIN_MODE=debug

# File Upload Configuration
UPLOAD_PATH=./uploads
MAX_FILE_SIZE=10MB

# CORS Configuration
ALLOWED_ORIGINS=http://localhost:5173,http://localhost:3000

# Admin User (Created on first run)
ADMIN_EMAIL=admin@gcx.com
ADMIN_PASSWORD=admin123
ADMIN_NAME=GCX Admin
```

## 🗄️ Step 2: Database Setup

The Go backend will automatically create the necessary tables in your existing `gcx_website` database.

**If you want to use a separate database for the CMS:**

1. Create a new database:
```sql
CREATE DATABASE gcx_cms CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

2. Update `DB_NAME=gcx_cms` in your `.env` file

## 🔑 Step 3: MySQL Password

If your MySQL root user has no password (default XAMPP/WAMP setup):
- Leave `DB_PASSWORD=` empty in `.env`

If your MySQL root user has a password:
- Set `DB_PASSWORD=your_actual_password` in `.env`

## 🚀 Step 4: Run the Backend

```bash
go run main.go
```

## ✅ Success Indicators

You should see:
```
✅ Connected to mysql database
✅ Created default admin user: admin@gcx.com
🚀 GCX CMS API Server starting on port 8080
```
