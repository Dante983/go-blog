-- Check if tables exist
SELECT table_name 
FROM information_schema.tables 
WHERE table_schema = 'public' 
AND table_name IN ('users', 'posts', 'categories', 'tags', 'post_tags')
ORDER BY table_name;

-- Check if admin user exists
SELECT username, email, is_admin, created_at 
FROM users 
WHERE username = 'admin'; 