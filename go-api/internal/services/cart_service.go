package services

import (
	"errors"

	"github.com/kunalavghade/Go-basics/go-api/internal/dto"
	"github.com/kunalavghade/Go-basics/go-api/internal/models"
	"gorm.io/gorm"
)

type CartService struct {
	db *gorm.DB
}

func NewCartService(db *gorm.DB) *CartService {
	return &CartService{
		db: db,
	}
}

func (s *CartService) GetCart(userID int) (*dto.CartResponse, error) {
	var cart models.Cart
	err := s.db.Preload("CartItems.Product.Category").Where("user_id = ?", userID).First(&cart).Error

	if err != nil {
		return nil, err
	}

	return s.convertToCartResponse(&cart)

}

func (s *CartService) AddToCart(userId int, req *dto.AddToCartRequest) (*dto.CartResponse, error) {
	var Product models.Product
	if err := s.db.First(&Product, req.ProductID).Error; err != nil {
		return nil, errors.New("Product Not Found")
	}
	if Product.Stock < req.Quantity {
		return nil, errors.New("Insuficient Stock")
	}
	var cart models.Cart
	if err := s.db.Where("user_id = ?", userId).First(&cart).Error; err != nil {
		cart = models.Cart{
			UserID: userId,
		}
		if err := s.db.Create(&cart).Error; err != nil {
			return nil, errors.New("Failed to Create Cart")
		}
	}

	var cartItems models.CartItem
	if err := s.db.Where("cart_id = ? AND product_id = ?", cart.ID, req.ProductID).First(&cartItems).Error; err != nil {
		cartItems = models.CartItem{
			CartID:    cart.ID,
			ProductID: req.ProductID,
			Quantity:  req.Quantity,
		}
		if err := s.db.Create(&cartItems).Error; err != nil {
			return nil, errors.New("Failed to Create Cart Item")
		}
	} else {
		cartItems.Quantity += req.Quantity
		if err := s.db.Save(&cartItems).Error; err != nil {
			return nil, errors.New("Failed to Update Cart Item")
		}
	}
	return s.GetCart(userId)
}

func (s *CartService) UpdateCart(userId int, req *dto.UpdateCartRequest) (*dto.CartResponse, error) {
	var cart models.Cart
	if err := s.db.Where("user_id = ?", userId).First(&cart).Error; err != nil {
		return nil, errors.New("Cart Not Found")
	}

	var cartItems models.CartItem
	if err := s.db.Where("cart_id = ? AND product_id = ?", cart.ID, req.ProductID).First(&cartItems).Error; err != nil {
		return nil, errors.New("Cart Item Not Found")
	}

	if req.Quantity <= 0 {
		if err := s.db.Delete(&cartItems).Error; err != nil {
			return nil, errors.New("Failed to Delete Cart Item")
		}
	} else {
		cartItems.Quantity = req.Quantity
		if err := s.db.Save(&cartItems).Error; err != nil {
			return nil, errors.New("Failed to Update Cart Item")
		}
	}
	return s.GetCart(userId)
}

func (s *CartService) RemoveFromCart(userId int, req *dto.RemoveFromCartRequest) (*dto.CartResponse, error) {
	var cart models.Cart
	if err := s.db.Where("user_id = ?", userId).First(&cart).Error; err != nil {
		return nil, errors.New("Cart Not Found")
	}

	var cartItems models.CartItem
	if err := s.db.Where("cart_id = ? AND product_id = ?", cart.ID, req.ProductID).First(&cartItems).Error; err != nil {
		return nil, errors.New("Cart Item Not Found")
	}

	if err := s.db.Delete(&cartItems).Error; err != nil {
		return nil, errors.New("Failed to Delete Cart Item")
	}
	return s.GetCart(userId)
}

func (s *CartService) convertToCartResponse(cart *models.Cart) (*dto.CartResponse, error) {

	cartItems := make([]dto.CartItemResponse, len(cart.CartItems))
	var total float64

	for i := range cart.CartItems {
		subtotal := float64(cart.CartItems[i].Product.Price * float64(cart.CartItems[i].Quantity))
		total += subtotal

		cartItems[i] = dto.CartItemResponse{
			ID: cart.CartItems[i].ID,
			Product: dto.ProductResponse{
				ID:          cart.CartItems[i].Product.ID,
				Name:        cart.CartItems[i].Product.Name,
				Price:       cart.CartItems[i].Product.Price,
				Stock:       cart.CartItems[i].Product.Stock,
				Description: cart.CartItems[i].Product.Description,
				CategoryID:  cart.CartItems[i].Product.CategoryID,
				SKU:         cart.CartItems[i].Product.SKU,
				Category: dto.CategoryResponse{
					ID:          cart.CartItems[i].Product.Category.ID,
					Name:        cart.CartItems[i].Product.Category.Name,
					Description: cart.CartItems[i].Product.Category.Description,
					IsActive:    cart.CartItems[i].Product.Category.IsActive,
				},
			},
			Quantity: cart.CartItems[i].Quantity,
			SubTotal: subtotal,
		}
	}

	return &dto.CartResponse{
		ID:       cart.ID,
		UserID:   cart.UserID,
		Total:    total,
		CartItem: cartItems,
	}, nil
}
