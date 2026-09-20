package dto

type AddToCartRequest struct {
	ProductID int `json:"product_id" binding:"required"`
	Quantity  int `json:"quantity" binding:"required"`
}

type UpdateCartRequest struct {
	Quantity int `json:"quantity" binding:"required"`
}

type CartResponse struct {
	ID       int                `json:"id"`
	UserID   int                `json:"user_id"`
	CartItem []CartItemResponse `json:"cart_item"`
	Total    float64            `json:"total"`
}

type CartItemResponse struct {
	ID       int             `json:"id"`
	Product  ProductResponse `json:"product"`
	Quantity int             `json:"quantity"`
	SubTotal float64         `json:"sub_total"`
}

type OrderResponse struct {
	ID          int                 `json:"id"`
	UserID      int                 `json:"user_id"`
	Status      string              `json:"status"`
	TotalAmount float64             `json:"total_amount"`
	OrderItems  []OrderItemResponse `json:"order_items"`
}

type OrderItemResponse struct {
	ID       int             `json:"id"`
	Product  ProductResponse `json:"product"`
	Quantity int             `json:"quantity"`
	Price    float64         `json:"price"`
}
