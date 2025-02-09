CREATE TABLE products (
                      "id" SERIAL PRIMARY KEY,
                      "brand_id" int  NOT NULL,
                      "category_id" int  NOT NULL,
                      "name" VARCHAR(255) NOT NULL,
                      "description" TEXT,
                      "price" DECIMAL NOT NULL,
                      "created_at" TIMESTAMP DEFAULT NOW(),
                      "updated_at" TIMESTAMP DEFAULT NOW(),
                      CONSTRAINT fk_brand FOREIGN KEY ("brand_id") REFERENCES brands(id) ON DELETE CASCADE,
                      CONSTRAINT fk_category FOREIGN KEY ("category_id") REFERENCES categories(id) ON DELETE CASCADE
);
