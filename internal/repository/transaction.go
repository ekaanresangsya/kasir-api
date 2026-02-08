package repository

import (
	"database/sql"
	"fmt"
	"kasir-api/internal/model"
	"log"
	"time"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) CreateTransaction(req model.CheckoutRequest) (*model.Transaction, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	totalAmount := 0
	details := make([]model.TransactionDetail, 0)

	for _, item := range req.Items {

		var name string
		var price int
		var stock *int

		// get data product
		querySelect := `SELECT name, price, stock
						FROM products
						WHERE id = $1`
		err := tx.QueryRow(querySelect, item.ProductID).
			Scan(&name, &price, &stock)
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product id %d not found", item.ProductID)
		}
		if err != nil {
			log.Printf("error get data product: %v", err)
			return nil, err
		}

		if stock != nil && *stock < item.Quantity {
			return nil, fmt.Errorf("stock not enough")
		}

		subtotal := item.Quantity * price
		totalAmount += subtotal

		// update stock
		queryUpdate := `UPDATE products
						SET stock = stock - $1
						WHERE id = $2`
		_, err = tx.Exec(queryUpdate, item.Quantity, item.ProductID)
		if err != nil {
			log.Printf("error update stock: %v", err)
			return nil, err
		}

		details = append(details, model.TransactionDetail{
			ProductID:   item.ProductID,
			ProductName: name,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		})
	}

	var transactionID int
	var createdAt time.Time
	// insert transaction
	queryInsert := `INSERT INTO transactions (total_amount)
				VALUES ($1)
				RETURNING id, created_at`
	err = tx.QueryRow(queryInsert, totalAmount).
		Scan(&transactionID, &createdAt)
	if err != nil {
		log.Printf("error insert transaction: %v", err)
		return nil, err
	}

	// insert transaction detail
	queryInsDetail := `INSERT INTO transaction_details (transaction_id, product_id, quantity, sub_total)
						VALUES `
	for i, detail := range details {
		details[i].TransactionID = transactionID
		queryInsDetail += fmt.Sprintf(`(%d, %d, %d, %d),`, transactionID, detail.ProductID, detail.Quantity, detail.Subtotal)
	}
	// remove last comma
	queryInsDetail = queryInsDetail[:len(queryInsDetail)-1]
	queryInsDetail += ` RETURNING id`
	rows, err := tx.Query(queryInsDetail)
	if err != nil {
		log.Printf("error insert transaction detail: %v", err)
		return nil, err
	}
	defer rows.Close()
	for i := 0; rows.Next(); i++ {
		err = rows.Scan(&details[i].ID)
		if err != nil {
			log.Printf("error scan transaction detail: %v", err)
			return nil, err
		}
	}

	return &model.Transaction{
		ID:          transactionID,
		TotalAmount: totalAmount,
		CreatedAt:   createdAt,
		Details:     details,
	}, nil
}
