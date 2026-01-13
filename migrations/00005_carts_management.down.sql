DROP INDEX IF EXISTS idx_cart_items_product_id;
DROP INDEX IF EXISTS idx_cart_items_cart_id;

DROP TABLE IF EXISTS cart_items;
DROP TABLE IF EXISTS carts;

DO $$
BEGIN
    IF EXISTS (
        SELECT 1 FROM pg_type WHERE typname = 'status_type'
    ) THEN
        DROP TYPE status_type;
    END IF;
END $$;
