package cart_queries

const (
	CREATECART=`
	INSERT INTO carts (user_id)
	VALUES ($1)
	RETURNING id;
	`
	GETUSERCART =`
	SELECT 
		c.id,
		c.user_id,
		c.status,
		COUNT(DISTINCT ci.product_id) AS item_count,
		COALESCE(
			JSON_AGG(
				JSON_BUILD_OBJECT(
					'product_id', ci.product_id,
					'quantity', ci.quantity,
					'snapshot_price', ci.unit_price,
					'product_details', JSON_BUILD_OBJECT(
						'name', p.name,
						'color', p.color,
						'origin', p.origin,
						'about', p.about
					)
				)
			) FILTER (WHERE ci.id IS NOT NULL),
			 '[]'::json
		) AS items
		FROM carts c 
		LEFT JOIN cart_items ci ON ci.cart_id = c.id 
		LEFT JOIN products p ON p.id = ci.product_id
		WHERE c.user_id = $1
		GROUP BY c.id
	`
	ADDTOCART=`
	INSERT INTO cart_items (cart_id, product_id, quantity, unit_price)
	VALUES ($1, $2, $3, $4)
	ON CONFLICT (cart_id, product_id) DO NOTHING;
	`

	UPDATEPRODUCTQUANTITY=`
	UPDATE cart_items
	SET quantity = $1
	WHERE cart_id = $2 AND product_id = $3;
	`

	DELETEFROMCART=`
	DELETE FROM cart_items
	WHERE cart_id = $1 AND product_id = $2;
	`

	CHECKAVAILABLECART=`
	SELECT
    c.id,
    COUNT(DISTINCT ci.product_id) AS item_count
	FROM carts c
	LEFT JOIN cart_items ci ON c.id = ci.cart_id
	WHERE c.user_id = $1
  	AND c.status = 'active'
	GROUP BY c.id;
	`
)