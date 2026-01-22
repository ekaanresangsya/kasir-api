package model

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}
