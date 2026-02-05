package repositories

import (
	"crud-category/models"
	"database/sql"
	"fmt"
	"time"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

func (repo *TransactionRepository) CreateTransaction(items []models.CheckoutItem) (*models.Transaction, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	totalAmount := 0
	details := []models.TransactionDetail{}

	for _, item := range items {
		var productPrice, stock int
		var productName string

		err := tx.QueryRow("SELECT price, stock, name FROM products WHERE id = $1", item.ProductID).Scan(&productPrice, &stock, &productName)
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product id %d not found", item.ProductID)
		}
		if err != nil {
			return nil, err
		}

		if stock < item.Quantity {
			return nil, fmt.Errorf("insufficient stock for product id %d", item.ProductID)
		}

		subtotal := productPrice * item.Quantity
		totalAmount += subtotal

		details = append(details, models.TransactionDetail{
			ProductID:   item.ProductID,
			ProductName: productName,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		})
	}

	var transactionID int
	err = tx.QueryRow("INSERT INTO transactions (total_amount) VALUES ($1) RETURNING id", totalAmount).Scan(&transactionID)
	if err != nil {
		return nil, err
	}

	for i := range details {
		_, err = tx.Exec("UPDATE products SET stock = stock - $1 WHERE id = $2", details[i].Quantity, details[i].ProductID)
		if err != nil {
			return nil, err
		}

		_, err = tx.Exec("INSERT INTO transaction_details (transaction_id, product_id, product_name, quantity, subtotal) VALUES ($1, $2, $3, $4, $5)",
			transactionID, details[i].ProductID, details[i].ProductName, details[i].Quantity, details[i].Subtotal)
		if err != nil {
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &models.Transaction{
		ID:          transactionID,
		Details:     details,
		TotalAmount: totalAmount,
		CreatedAt:   time.Now(),
	}, nil
}

func (r *TransactionRepository) TransactionReport(start, end string) (*models.TransactionReport, error) {
	query := `	WITH FilteredTransactions AS (
			SELECT
				id, total_amount, created_at
			FROM transactions
			WHERE created_at >= $1 AND created_at < ($2::date + INTERVAL '1 day')
		),
		TransactionSummary AS (
			SELECT
				COUNT(id) AS total_transactions,
				SUM(total_amount) AS total_revenue
			FROM FilteredTransactions
		),
		ProductSales AS (
			SELECT
				p.id AS product_id,
				p.name AS product_name,
				SUM(td.quantity) AS sold_quantity
			FROM transaction_details td
			JOIN products p ON td.product_id = p.id
			-- Join with FilteredTransactions to ensure we only count sales
			-- from transactions that fall within our date range.
			JOIN FilteredTransactions ft ON td.transaction_id = ft.id
			GROUP BY p.id, p.name
		),
		RankedProductSales AS (
			SELECT
				product_id,
				product_name,
				sold_quantity,
				ROW_NUMBER() OVER (ORDER BY sold_quantity DESC) as rn
			FROM ProductSales
		)
		SELECT
			COALESCE(ts.total_transactions, 0) AS total_transactions,
			COALESCE(ts.total_revenue, 0) AS total_revenue,
			rps.product_id AS best_seller_id,
			rps.product_name AS best_seller_name,
			rps.sold_quantity AS best_seller_sold_quantity
		FROM TransactionSummary ts
		LEFT JOIN RankedProductSales rps ON rps.rn = 1;`
	var report models.TransactionReport
	var totalRevenue sql.NullInt64
	var totalTransactions sql.NullInt64
	var bestSellerID sql.NullInt64
	var bestSellerName sql.NullString
	var bestSellerSoldQuantity sql.NullInt64

	err := r.db.QueryRow(query, start, end).Scan(
		&totalTransactions,
		&totalRevenue,
		&bestSellerID,
		&bestSellerName,
		&bestSellerSoldQuantity,
	)

	if err == sql.ErrNoRows {
		return &models.TransactionReport{
			TotalRevenue:      0,
			TotalTransactions: 0,
			BestSeller:        models.BestSellerProduct{},
		}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error fetching transaction report: %w", err)
	}

	if totalTransactions.Valid {
		report.TotalTransactions = int(totalTransactions.Int64)
	} else {
		report.TotalTransactions = 0
	}

	if totalRevenue.Valid {
		report.TotalRevenue = int(totalRevenue.Int64)
	} else {
		report.TotalRevenue = 0
	}

	if bestSellerID.Valid && bestSellerName.Valid && bestSellerSoldQuantity.Valid {
		report.BestSeller = models.BestSellerProduct{
			ID:           int(bestSellerID.Int64),
			Name:         bestSellerName.String,
			SoldQuantity: int(bestSellerSoldQuantity.Int64),
		}
	} else {
		report.BestSeller = models.BestSellerProduct{}
	}

	return &report, nil
}
