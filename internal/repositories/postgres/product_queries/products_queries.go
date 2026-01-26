package product_queries

const (
	INSERTMASTERPRODUCT=`
	INSERT INTO products (
	name, origin, about, image_urls
	)
	VALUES ($1, $2, $3, $4)
	RETURNING id
	`
	ADDPRODUCTVARIANT=`
	INSERT INTO product_variants (
	product_id, color, quantity, weight, price, image_urls
	)
	VALUES($1, $2, $3, $4, $5, $6)
	`
	GETPRODUCTS = `
	SELECT
    	p.id,
    	p.name,
    	p.origin,
    	p.about,
    	p.image_urls,
    	p.created_at,
    	COALESCE(v_agg.total_qty, 0) AS total_quantity,
    	COALESCE(c_agg.list, '[]'::json) AS categories,
    	COALESCE(v_agg.list, '[]'::json) AS variants
	FROM products p
	LEFT JOIN LATERAL (
    	SELECT JSON_AGG(JSON_BUILD_OBJECT(
        	'id', c.id,
        	'name', c.name,
        	'description', c.description,
        	'image_url', c.image_url
    	)) AS list
    FROM product_category pc
    JOIN category c ON pc.category_id = c.id
    WHERE pc.product_id = p.id
	) c_agg ON TRUE
	LEFT JOIN LATERAL (
    	SELECT 
        	JSON_AGG(JSON_BUILD_OBJECT(
            	'id', v.id,
            	'color', v.color,
            	'quantity', v.quantity,
            	'price', v.price,
            	'image_urls', v.image_urls,
            	'weight', v.weight,
            	'created_at', v.created_at
        	)) AS list,
        SUM(v.quantity) AS total_qty,
        MIN(v.price) AS min_price, -- Useful for filtering
        MAX(v.price) AS max_price
    FROM product_variants v
    WHERE v.product_id = p.id
    AND ($4::TEXT IS NULL OR v.color ILIKE '%' || $4 || '%')
	) v_agg ON TRUE
	WHERE 
    	($1::TEXT IS NULL OR p.name ILIKE '%' || $1 || '%')
    	AND ($5::TEXT IS NULL OR p.origin ILIKE '%' || $5 || '%')
    	AND ($2::BIGINT IS NULL OR v_agg.max_price >= $2)
    	AND ($3::BIGINT IS NULL OR v_agg.min_price <= $3)
    	AND ($4::TEXT IS NULL OR v_agg.list IS NOT NULL)
	ORDER BY p.created_at DESC
	LIMIT $6 OFFSET $7;
	`

	GETPRODUCTBYID=`
	WITH product_base AS (
    SELECT * FROM products WHERE id = $1
	)
	SELECT
		pb.id,
		pb.name,
		pb.origin,
		pb.about,
		pb.image_urls,
		pb.created_at,
		COALESCE(
        (SELECT SUM(v.quantity) FROM product_variants v WHERE v.product_id = pb.id), 0
    	) AS total_quantity,
		COALESCE(
        	(SELECT JSON_AGG(JSON_BUILD_OBJECT(
            	'id', c.id,
            	'name', c.name,
            	'description', c.description
        )	)
         FROM product_category pc
         JOIN category c ON pc.category_id = c.id
         WHERE pc.product_id = pb.id
        ), '[]'
    ) AS categories,
		COALESCE(
        	(SELECT JSON_AGG(JSON_BUILD_OBJECT(
            	'id', v.id,
				'product_id', v.product_id,
            	'color', v.color,
            	'quantity', v.quantity,
				'price', v.price,
				'image_urls', v.image_urls,
            	'weight', v.weight,
				'created_at', v.created_at,
				'updated_at', v.updated_at
        )	)
         FROM product_variants v
         WHERE v.product_id = pb.id
        ), '[]'
    	) AS variants
	FROM product_base pb
	`
	CREATEPRODUCTCATEGORY=`
	INSERT INTO category (
	name, description, image_url
	)
	VALUES ($1, $2, $3)
	RETURNING id, name, description, image_url
	`
	DELETEPRODUCTCATEGORY=`
	DELETE FROM category
	WHERE id = $1
	`
	UPDATEPRODUCTCATEGORY=`
	UPDATE category
	SET 
	    name = COALESCE($1, name),
		description = COALESCE($2, description),
		image_url = COALESCE($3, image_url)
	WHERE id = $4
	RETURNING id, name, description, image_url
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
	WITH RECURSIVE category_tree AS (
    SELECT id FROM category WHERE id = $1
    UNION ALL
    SELECT c.id FROM category c
    JOIN category_tree ct ON c.parent_id = ct.id
	),
	matching_products AS (
    	SELECT DISTINCT pc.product_id
    	FROM product_category pc
    	JOIN category_tree ct ON ct.id = pc.category_id
	)
	SELECT 
    	p.id,
    	p.name,
    	p.origin,
    	p.about,
    	p.image_urls,
    	p.created_at,
    	COALESCE(v_agg.total_quantity, 0) AS total_quantity,
    	COALESCE(v_agg.list, '[]'::json) AS variants
	FROM matching_products mp
	JOIN products p ON p.id = mp.product_id
	LEFT JOIN LATERAL (
    	SELECT 
        	SUM(v.quantity) AS total_quantity,
        	MIN(v.price) AS min_price,
        	JSON_AGG(JSON_BUILD_OBJECT(
            	'id', v.id,
            	'color', v.color,
            	'quantity', v.quantity,
            	'price', v.price,
				'weight', v.weight
        	)) AS list
    	FROM product_variants v
    	WHERE v.product_id = p.id
	) v_agg ON TRUE
	ORDER BY p.created_at DESC
	LIMIT $2 OFFSET $3;
	`

	GETCATEGORY=`
	SELECT id, name, description, image_url
	FROM category
	WHERE id = $1
	`
	GETSUBCATEGORIES=`
	SELECT id, name, description, image_url
	FROM category 
	WHERE parent_id = $1
	`
)