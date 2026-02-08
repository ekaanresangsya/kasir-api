package repository

import (
	"database/sql"
	"kasir-api/internal/model"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetAll() ([]model.Category, error) {
	query := `SELECT id, name, description
				FROM categories`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []model.Category
	for rows.Next() {
		var category model.Category
		if err := rows.Scan(&category.ID, &category.Name, &category.Description); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (r *CategoryRepository) GetByID(id int64) (*model.Category, error) {
	query := `SELECT id, name, description
				FROM categories
				WHERE id = $1`

	row := r.db.QueryRow(query, id)
	var category model.Category
	if err := row.Scan(&category.ID, &category.Name, &category.Description); err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *CategoryRepository) Create(category *model.Category) (*model.Category, error) {
	query := `INSERT INTO categories (name, description)
				VALUES ($1, $2)
				RETURNING id`

	row := r.db.QueryRow(query, category.Name, category.Description)
	if err := row.Scan(&category.ID); err != nil {
		return nil, err
	}

	return category, nil
}

func (r *CategoryRepository) Update(id int64, category *model.Category) error {
	query := `UPDATE categories
				SET name = $1, description = $2
				WHERE id = $3`

	result, err := r.db.Exec(query, category.Name, category.Description, id)
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

func (r *CategoryRepository) Delete(id int64) error {
	query := `DELETE FROM categories
				WHERE id = $1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
