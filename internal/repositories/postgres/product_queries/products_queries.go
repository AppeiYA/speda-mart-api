package product_queries

const (
	GETPRODUCTS = `
	SELECT
		id,
		name,
		quantity,
		color,
		price,
		origin,
		about,
		image_urls,
		created_at
	FROM products
	WHERE 
		($1::TEXT IS NULL OR name ILIKE '%' || $1 || '%')
		AND ($2::BIGINT IS NULL OR price >= $2)
		AND ($3::BIGINT IS NULL OR price <= $3)
		AND ($4::TEXT IS NULL OR color ILIKE '%' || $4 || '%')
		AND ($5::TEXT IS NULL OR origin ILIKE '%' || $5 || '%')
	ORDER BY created_at DESC
	LIMIT $6
	OFFSET $7;
	`

	GETPRODUCTBYID=`
	SELECT
		id,
		name,
		quantity,
		color,
		price,
		origin,
		about,
		image_urls,
		created_at
	FROM products
	WHERE id = $1
	`

)