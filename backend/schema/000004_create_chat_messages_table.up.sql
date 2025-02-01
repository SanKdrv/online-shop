CREATE TABLE chat_messages (
                               "id" SERIAL PRIMARY KEY,
                               "order_id" int  NOT NULL,
                               "user_id" int  NOT NULL,
                               "message" TEXT NOT NULL,
                               "created_at" TIMESTAMP DEFAULT NOW(),
                               CONSTRAINT fk_chat_owner FOREIGN KEY ("chat_owner_id") REFERENCES users(id) ON DELETE CASCADE
);
