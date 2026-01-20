package seed

import (
	"e-commerce/internal/db"
	"e-commerce/internal/utils"
)

func SeedAdmin(db *db.DB, adminKey string, adminEmail string) error {
	// Check if admin already exists
    var count int
    err := db.Get(&count, "SELECT COUNT(*) FROM users WHERE role='admin'")
    if err != nil {
        return err
    }
    if count > 0 {
        return nil 
    }

    // Generate bcrypt hash for password
    hashedPassword, err := utils.HashPassword(adminKey)
    if err != nil {
        return err
    }

    // Insert admin user
    _, err = db.Exec(`
        INSERT INTO users (first_name, last_name, email, hash, role)
        VALUES ('Admin', 'User', $1, $2, 'admin')
    `,adminEmail, string(hashedPassword))

    return err
}