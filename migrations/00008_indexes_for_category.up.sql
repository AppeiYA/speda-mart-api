CREATE INDEX IF NOT EXISTS idx_category_parent_id ON category(parent_id);
CREATE INDEX IF NOT EXISTS idx_product_category_category_id ON product_category(category_id);
CREATE INDEX IF NOT EXISTS idx_product_category_product_id ON product_category(product_id);
