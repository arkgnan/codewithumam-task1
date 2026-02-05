package database

import (
	"database/sql"
	"log"
)

func runMigrations(db *sql.DB) error {
	// SQL statements for creating tables
	// Using IF NOT EXISTS ensures that these statements can be run multiple times without error.
	migrationSQL := `
	CREATE TABLE IF NOT EXISTS categories (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		description TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		price INT NOT NULL,
		stock INT NOT NULL,
		category_id INT REFERENCES categories(id) ON DELETE SET NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS transactions (
		id SERIAL PRIMARY KEY,
		total_amount INT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS transaction_details (
		id SERIAL PRIMARY KEY,
		transaction_id INT REFERENCES transactions(id) ON DELETE CASCADE,
		product_id INT REFERENCES products(id) ON DELETE SET NULL,
		product_name VARCHAR(255) NOT NULL,
		quantity INT NOT NULL,
		subtotal INT NOT NULL
	);
	`

	// Execute the SQL statements
	// db.Exec is suitable for statements that don't return rows, like CREATE TABLE.
	_, err := db.Exec(migrationSQL)
	if err != nil {
		log.Printf("Error executing migrations: %v", err)
		return err
	}

	return nil
}
