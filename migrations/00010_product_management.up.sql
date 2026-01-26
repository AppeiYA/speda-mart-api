CREATE TABLE IF NOT EXISTS product_variants (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    product_id UUID NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    color TEXT NOT NULL,
    quantity INT DEFAULT 0,
    weight BIGINT,
    price BIGINT NOT NULL,
    image_urls JSONB NOT NULL DEFAULT '[]', 
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

INSERT INTO product_variants (product_id, color, quantity, price, image_urls, created_at, updated_at)
SELECT id, color, quantity, price, image_urls, created_at, updated_at
FROM products;

ALTER TABLE IF EXISTS products 
    DROP COLUMN color,
    DROP COLUMN price,
    DROP COLUMN quantity;