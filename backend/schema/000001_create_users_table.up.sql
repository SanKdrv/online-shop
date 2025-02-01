CREATE TABLE if not exists users(
                      id SERIAL PRIMARY KEY,
                      username VARCHAR(255) NOT NULL UNIQUE,
                      name VARCHAR(255),
                      google_name VARCHAR(255),
                      email VARCHAR(255) NOT NULL UNIQUE,
                      password_hash VARCHAR(255) NOT NULL,
                      role VARCHAR(255) DEFAULT 'user',
                      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                      updated_at TIMESTAMP
)