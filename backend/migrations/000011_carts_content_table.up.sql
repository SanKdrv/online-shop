CREATE TABLE carts_content (
                       "id" SERIAL PRIMARY KEY,
                       "user_id" INT NOT NULL,
                       "product_id" INT NOT NULL,
                       "count" INT NOT NULL,
                       "created_at" TIMESTAMP DEFAULT NOW(),
                       "updated_at" TIMESTAMP DEFAULT NOW(),
                       CONSTRAINT fk_user FOREIGN KEY ("user_id") REFERENCES users(id) ON DELETE CASCADE,
                       CONSTRAINT fk_product FOREIGN KEY ("product_id") REFERENCES products(id) ON DELETE CASCADE
);