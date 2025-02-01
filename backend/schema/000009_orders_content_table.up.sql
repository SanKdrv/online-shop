CREATE TABLE orders_content (
                        "id" SERIAL PRIMARY KEY,
                        "count" DECIMAL NOT NULL,
                        CONSTRAINT fk_order FOREIGN KEY ("user_id") REFERENCES orders(id) ON DELETE CASCADE,
                        CONSTRAINT fk_product FOREIGN KEY ("product_id") REFERENCES products(id) ON DELETE CASCADE,
);
