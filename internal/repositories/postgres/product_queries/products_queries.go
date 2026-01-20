package product_queries

const (
	GETPRODUCTS = `
	SELECT
		p.id,
		p.name,
		p.quantity,
		p.color,
		p.price,
		COALESCE(
			JSON_AGG(
				JSON_BUILD_OBJECT(
					'name', c.name,
					'description', c.description
				)
			) FILTER (WHERE c.id IS NOT NULL), '[]'
		) AS categories,
		p.origin,
		p.about,
		p.image_urls,
		p.created_at
	FROM products p
	LEFT JOIN product_category pc ON pc.product_id = p.id
	LEFT JOIN category c ON pc.category_id = c.id
	WHERE 
		($1::TEXT IS NULL OR p.name ILIKE '%' || $1 || '%')
		AND ($2::BIGINT IS NULL OR p.price >= $2)
		AND ($3::BIGINT IS NULL OR p.price <= $3)
		AND ($4::TEXT IS NULL OR p.color ILIKE '%' || $4 || '%')
		AND ($5::TEXT IS NULL OR p.origin ILIKE '%' || $5 || '%')
	GROUP BY
    p.id, p.name, p.quantity, p.color, p.price, p.origin, p.about, p.image_urls, p.created_at
	ORDER BY p.created_at DESC
	LIMIT $6
	OFFSET $7;
	`

	GETPRODUCTBYID=`
	SELECT
		p.id,
		p.name,
		p.quantity,
		p.color,
		p.price,
		COALESCE(
		    JSON_AGG(
				JSON_BUILD_OBJECT(
				    'name', c.name,
					'description', c.description
				)
			) FILTER (WHERE c.id IS NOT NULL), '[]'
		) AS categories,
		p.origin,
		p.about,
		p.image_urls,
		p.created_at
	FROM products p
	LEFT JOIN product_category pc ON pc.product_id = p.id
	LEFT JOIN category c ON pc.category_id = c.id
	WHERE p.id = $1
	GROUP BY p.id, p.name, p.quantity, p.color, p.price, p.origin, p.about, p.image_urls, p.created_at
	`
	CREATEPRODUCTCATEGORY=`
	INSERT INTO category (
	name, description
	)
	VALUES ($1, $2)
	RETURNING id, name, description
	`
	DELETEPRODUCTCATEGORY=`
	DELETE FROM category
	WHERE id = $1
	`
	UPDATEPRODUCTCATEGORY=`
	UPDATE category
	SET 
	    name = COALESCE($1, name),
		description = COALESCE($2, description)
	WHERE id = $3
	RETURNING id, name, description
	`
	ADDPRODUCTTOCATEGORY=`
	INSERT INTO product_category (
	product_id, category_id
	)
	VALUES ($1, $2)
	ON CONFLICT (product_id, category_id) DO NOTHING;
	`
	REMOVEPRODUCTFROMCATEGORY=`
	DELETE FROM product_category
	WHERE product_id = $1 AND category_id = $2
	`
	GETPRODUCTSINCATEGORY=`
	SELECT 
		p.id,
		p.name,
		p.quantity,
		p.image_urls,
		p.color,
		p.price,
		p.origin,
		p.about
	FROM products p
	JOIN product_category pc ON p.id = pc.product_id
	WHERE pc.category_id = $1
	LIMIT $2
	OFFSET $3
	`

	GETCATEGORY=`
	SELECT name, description
	FROM category
	WHERE id = $1
	`

)