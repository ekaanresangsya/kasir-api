package repository

import (
	"database/sql"
	"kasir-api/internal/model"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetAll(req model.GetProductReq) ([]model.Product, error) {
	query := `SELECT p.id, p.name, p.stock, p.price, p.category_id, c.id, c.name, c.description
				FROM products p
				JOIN categories c ON p.category_id = c.id`

	args := []any{}
	if req.Name != "" {
		query += " WHERE p.name ILIKE $1"
		args = append(args, "%"+req.Name+"%")
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Stock, &product.Price,
			&product.CategoryID, &product.Category.ID, &product.Category.Name, &product.Category.Description); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *ProductRepository) GetByID(id int) (*model.Product, error) {
	query := `SELECT p.id, p.name, p.stock, p.price, p.category_id, c.id, c.name, c.description
				FROM products p
				JOIN categories c ON p.category_id = c.id
				WHERE p.id = $1`

	row := r.db.QueryRow(query, id)
	var product model.Product
	if err := row.Scan(&product.ID, &product.Name, &product.Stock, &product.Price, &product.CategoryID,
		&product.Category.ID, &product.Category.Name, &product.Category.Description); err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) Create(product *model.Product) error {
	query := `INSERT INTO products (name, stock, price, category_id)
				VALUES ($1, $2, $3, $4)
				RETURNING id`

	row := r.db.QueryRow(query, product.Name, product.Stock, product.Price, product.CategoryID)
	if err := row.Scan(&product.ID); err != nil {
		return err
	}

	return nil
}

func (r *ProductRepository) Update(product *model.Product) error {
	query := `UPDATE products
				SET name = $1, stock = $2, price = $3, category_id = $4
				WHERE id = $5`

	result, err := r.db.Exec(query, product.Name, product.Stock, product.Price, product.CategoryID, product.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *ProductRepository) Delete(id int64) error {
	query := `DELETE FROM products
				WHERE id = $1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
