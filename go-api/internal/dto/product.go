package dto

type CreateCategoryRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	IsActive    bool   `json:"is_active" binding:"required"`
}

type UpdateCategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}

type CategoryResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}

type CreateProductRequest struct {
	CategoryID  int     `json:"category_id" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Stock       int     `json:"stock" binding:"required"`
	SKU         string  `json:"sku" binding:"required"`
}

type UpdateProductRequest struct {
	CategoryID  int     `json:"category_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	SKU         string  `json:"sku"`
}

type ProductResponse struct {
	ID          int     `json:"id"`
	CategoryID  int     `json:"category_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	SKU         string  `json:"sku"`
	Category    CategoryResponse
	Images      []ProductImageResponse
}

type ProductImageResponse struct {
	ID        int    `json:"id"`
	ProductID int    `json:"product_id"`
	ImageURL  string `json:"image_url"`
}
