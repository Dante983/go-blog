-- Create users table for authentication
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    is_admin BOOLEAN DEFAULT FALSE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Insert default admin user (password: admin123)
-- This hash is for 'admin123' - CHANGE THIS IN PRODUCTION!
-- Generated with: go run scripts/hash_password.go admin123
INSERT INTO users (username, email, password_hash, is_admin) VALUES 
('admin', 'admin@example.com', '$2a$10$Sv5V1DhIz6b5FTO3c9DJOuV0s7RqGa5VDKkNvXxMjBLBVyHuBBCZW', TRUE); 