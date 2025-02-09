-- We'll use the pgcrypto extension for password hashing
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- Insert admin user with hashed password
-- Password: admin123
INSERT INTO users (id, email, password, name, created_at, updated_at)
VALUES (
    '00000000-0000-0000-0000-000000000001', -- fixed UUID for admin
    'admin@gmail.com',
    crypt('admin123', gen_salt('bf')), -- using Blowfish encryption (same as bcrypt)
    'Admin User',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
) ON CONFLICT (email) DO NOTHING;
