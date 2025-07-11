# Authentication Setup

## 🔐 Implemented Features

- ✅ Session-based authentication with secure cookies
- ✅ Bcrypt password hashing
- ✅ Login/logout flow with HTMX
- ✅ Protected admin routes with middleware
- ✅ Retro terminal-style login page

## 📦 Required Packages

Install the required packages:

```bash
go get github.com/gorilla/sessions
go get golang.org/x/crypto/bcrypt
```

## 🗄️ Database Setup

Apply the users table migration:

```bash
# For local PostgreSQL
psql -U postgres -d go_blog < migrations/003_create_users.sql

# For Supabase
# Use the SQL editor in Supabase Dashboard and run the migration scripts
```

This creates:
- `users` table with secure password storage
- Default admin user (username: `admin`, password: `admin123`)

## 🔑 Environment Variables

Add to your `.env` file:

```
# For Supabase
DB_HOST=your-project-db.supabase.co
DB_USER=postgres
DB_PASS=your-supabase-password
DB_NAME=postgres
DB_SSL_MODE=require

# Session key
SESSION_KEY=your-32-byte-long-secret-key-here!!
```

For production, generate a secure key:
```bash
openssl rand -base64 32
```

## 🛡️ Protected Routes

The following routes are now protected:
- `/admin/posts/new` - Requires admin authentication
- `/admin/posts` - Requires admin authentication

## 🚀 Usage

1. Navigate to `/login`
2. Enter credentials:
   - Username: `admin`
   - Password: `admin123`
3. You'll be redirected to the admin panel

## 🔧 Creating New Users

Use the provided script to hash passwords:

```bash
go run scripts/hash_password.go <password>
```

Then insert into database:
```sql
INSERT INTO users (username, email, password_hash, is_admin) 
VALUES ('username', 'email@example.com', 'hash_from_script', TRUE);
```

## 🔒 Security Notes

- Change the default admin password in production
- Use HTTPS in production (set `Secure: true` in session options)
- Generate a strong SESSION_KEY for production
- Consider adding CSRF protection for forms
- Add rate limiting for login attempts

## 🎨 UI Features

- Retro terminal login screen with ASCII art
- Dynamic navigation showing Login/Logout based on auth status
- HTMX-powered login with real-time feedback
- Terminal-style error and success messages 