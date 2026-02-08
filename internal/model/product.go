package model

type Product struct {
	ID         int64    `json:"id"`
	Name       string   `json:"name"`
	Stock      *int64   `json:"stock"`
	Price      int64    `json:"price"`
	CategoryID int64    `json:"category_id"`
	Category   Category `json:"category"`
}

type GetProductReq struct {
	Name string `form:"name"`
}

type CreateProductReq struct {
	Name       string `json:"name" binding:"required"`
	Stock      *int64 `json:"stock"`
	Price      int64  `json:"price" binding:"required"`
	CategoryID int64  `json:"category_id" binding:"required"`
}

type UpdateProductReq struct {
	Name       string `json:"name" binding:"required"`
	Stock      *int64 `json:"stock"`
	Price      int64  `json:"price" binding:"required"`
	CategoryID int64  `json:"category_id" binding:"required"`
}

type ProductResponse struct {
	ID       int64    `json:"id"`
	Name     string   `json:"name"`
	Stock    *int64   `json:"stock"`
	Price    int64    `json:"price"`
	Category Category `json:"category"`
}

func (p *Product) ToResponse() *ProductResponse {
	return &ProductResponse{
		ID:       p.ID,
		Name:     p.Name,
		Stock:    p.Stock,
		Price:    p.Price,
		Category: p.Category,
	}
}
