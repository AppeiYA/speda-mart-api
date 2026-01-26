-- 1. Add the columns back to products
ALTER TABLE IF EXISTS products 
    ADD COLUMN color TEXT,
    ADD COLUMN price BIGINT,
    ADD COLUMN quantity INT;

-- 2. Migrate data back from product_variants to products
UPDATE products p
SET 
    color = v.color,
    quantity = v.quantity,
    price = v.price
FROM (
    SELECT DISTINCT ON (product_id) product_id, color, quantity, price
    FROM product_variants
    ORDER BY product_id, created_at ASC
) v
WHERE p.id = v.product_id;

-- 3. Drop the variants table
DROP TABLE IF EXISTS product_variants;