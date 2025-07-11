# PostgreSQL Migration Guide

## 🔄 Migration from MySQL to PostgreSQL

This project has been migrated from MySQL to PostgreSQL to support Supabase hosting.

## 📋 What Changed

### 1. **Database Driver**
- Changed from `github.com/go-sql-driver/mysql` to `github.com/lib/pq`
- Updated connection string format in `internal/db/postgres.go`

### 2. **SQL Syntax Updates**
- **Placeholders**: Changed from `?` to `$1, $2, $3...`
- **Auto-increment**: Changed from `AUTO_INCREMENT` to `SERIAL`
- **Timestamps**: Changed from `DATETIME` to `TIMESTAMP`
- **Update triggers**: Added PostgreSQL triggers for `updated_at` columns

### 3. **Migration Files**
All migration files have been updated with PostgreSQL syntax:
- `001_init.sql` - Posts table
- `002_add_categories_tags.sql` - Categories and tags
- `003_create_users.sql` - Users table with triggers

## 🚀 Setup Instructions

### Local Development with Docker

1. **Update your `.env` file:**
```env
DB_HOST=postgres
DB_USER=bloguser
DB_PASS=yourpassword
DB_NAME=goblog
DB_SSL_MODE=disable
SESSION_KEY=your-32-byte-session-key-here
```

2. **Start the services:**
```bash
docker-compose down -v  # Remove old MySQL volumes
docker-compose up -d
```

### Production with Supabase

1. **Create a Supabase project** at https://supabase.com

2. **Get your database credentials** from:
   - Supabase Dashboard → Settings → Database

3. **Update your `.env` file:**
```env
DB_HOST=db.xxxxxxxxxxxxxxxxxxxx.supabase.co
DB_USER=postgres
DB_PASS=your-supabase-password
DB_NAME=postgres
DB_SSL_MODE=require
SESSION_KEY=your-32-byte-session-key-here
```

4. **Run migrations in Supabase:**
   - Go to SQL Editor in Supabase Dashboard
   - Run each migration file in order:
     1. `migrations/001_init.sql`
     2. `migrations/002_add_categories_tags.sql`
     3. `migrations/003_create_users.sql`

## 📝 Code Changes Summary

### Models
- `internal/models/post.go` - Updated queries to use `$1, $2` placeholders
- `internal/models/user.go` - Updated queries to use `$1, $2, $3, $4` placeholders

### Handlers
- `internal/handlers/admin.go` - Updated INSERT query placeholders

### Database Connection
- `internal/db/postgres.go` (renamed from mysql.go) - New PostgreSQL connection logic with SSL support

## 🔧 Troubleshooting

### SSL Connection Issues
- For local development: Use `DB_SSL_MODE=disable`
- For Supabase: Use `DB_SSL_MODE=require`

### Port Already in Use
```bash
# Stop existing PostgreSQL container
docker-compose down

# Or use a different port in docker-compose.yml
ports:
  - "5433:5432"  # Use 5433 on host
```

### Migration Errors
- Ensure you run migrations in order (001, 002, 003)
- Check that tables don't already exist
- For fresh start: Drop all tables and re-run migrations

## 🎯 Benefits of PostgreSQL

1. **Supabase Compatible** - Direct integration with Supabase
2. **Better JSON Support** - Native JSONB type for future features
3. **Full Text Search** - Built-in FTS capabilities
4. **Row Level Security** - Enhanced security features
5. **Better Performance** - For complex queries and concurrent users

## 🔐 Security Notes

- Always use SSL (`DB_SSL_MODE=require`) in production
- Rotate your `SESSION_KEY` regularly
- Use strong passwords for database users
- Enable Row Level Security (RLS) in Supabase for additional protection 