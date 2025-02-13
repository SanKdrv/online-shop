CREATE TABLE products_images
(
    "id"         SERIAL PRIMARY KEY,
    "image_order" INT DEFAULT 0,
    "product_id" INT NOT NULL,
    "image_hash" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP DEFAULT NOW(),
    "updated_at" TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_product FOREIGN KEY ("product_id") REFERENCES products(id) ON DELETE CASCADE
);