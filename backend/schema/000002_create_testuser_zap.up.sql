INSERT INTO users (username, name, google_name, email, password_hash, created_at, updated_at)
VALUES
    ('test_user', 'Test User', 'Test Google Name', 'test_user@example.com', 'hashed_password', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
