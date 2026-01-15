package user_queries

const (
	CREATEUSER = `
	INSERT INTO users (first_name, last_name, email, hash)
	VALUES ($1, $2, $3, $4)
	RETURNING id, first_name, last_name, email, hash, role, created_at, updated_at;
	`

	GETUSERBYEMAIL = `
	SELECT first_name, last_name, email, role, created_at
	FROM users 
	WHERE email = $1
	`
	GETUSERFORAUTH = ` SELECT * FROM users WHERE email = $1`
)
