CREATE TABLE orders_content (
                      "id" SERIAL PRIMARY KEY,
                      "order_id" INT NOT NULL,
                      "product_id" INT NOT NULL,
                      "count" INT NOT NULL,
                      "sum_price" DECIMAL NOT NULL,
                      CONSTRAINT fk_order FOREIGN KEY ("order_id") REFERENCES orders(id) ON DELETE CASCADE,
                      CONSTRAINT fk_product FOREIGN KEY ("product_id") REFERENCES products(id) ON DELETE CASCADE
);
