CREATE TABLE orders (
                          "id" SERIAL PRIMARY KEY,
                          "cost" DECIMAL NOT NULL,
                          "status" VARCHAR(255) NOT NULL,
                          "created_at" TIMESTAMP DEFAULT NOW(),
                          "updated_at" TIMESTAMP DEFAULT NOW(),
                          CONSTRAINT fk_user FOREIGN KEY ("user_id") REFERENCES users(id) ON DELETE CASCADE,
);
