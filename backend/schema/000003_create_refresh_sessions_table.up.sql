CREATE TABLE refresh_sessions (
                                 "id" SERIAL PRIMARY KEY,
                                 "user_id" BIGINT REFERENCES users(id) ON DELETE CASCADE,
                                 "refresh_token" uuid NOT NULL UNIQUE,
                                 "user_agent" VARCHAR(200) NOT NULL, /* user-agent */
                                 "expires_in" BIGINT NOT NULL,
                                 "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);