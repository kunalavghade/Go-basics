package services

import (
	"errors"
	"strconv"

	"github.com/kunalavghade/Go-basics/go-api/internal/dto"
	"github.com/kunalavghade/Go-basics/go-api/internal/interfaces"
	"github.com/kunalavghade/Go-basics/go-api/internal/models"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type OrderService struct {
	db        *gorm.DB
	publisher interfaces.EventPublisher
}

func NewOrderService(db *gorm.DB, publisher interfaces.EventPublisher) *OrderService {
	return &OrderService{
		db:        db,
		publisher: publisher,
	}
}

func (os *OrderService) CreateOrder(userID int) (*dto.OrderResponse, error) {
	var orderResponse *dto.OrderResponse

	err := os.db.Transaction(func(tx *gorm.DB) error {
		var cart models.Cart
		if err := tx.Preload("CartItems.Product.Category").Where("user_id = ?", userID).First(&cart).Error; err != nil {
			return errors.New("Cart Not Found")
		}
		if len(cart.CartItems) == 0 {
			return errors.New("Cart is empty")
		}

		var totalAmount float64
		for _, item := range cart.CartItems {
			if item.Product.Stock < item.Quantity {
				return errors.New("Product " + item.Product.Name + " is out of stock")
			}
			totalAmount += item.Product.Price * float64(item.Quantity)
		}

		order := models.Order{
			UserID:      userID,
			Status:      models.OrderStatusPending,
			TotalAmount: totalAmount,
		}
		if err := tx.Create(&order).Error; err != nil {
			return errors.New("Failed to create order")
		}

		orderItemResponses := make([]dto.OrderItemResponse, len(cart.CartItems))
		for i, item := range cart.CartItems {
			orderItem := models.OrderItem{
				OrderID:   order.ID,
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
				Price:     item.Product.Price,
			}
			if err := tx.Create(&orderItem).Error; err != nil {
				return errors.New("Failed to create order item")
			}

			item.Product.Stock -= item.Quantity
			if err := tx.Save(&item.Product).Error; err != nil {
				return errors.New("Failed to update stock")
			}

			orderItemResponses[i] = dto.OrderItemResponse{
				ID:       orderItem.ID,
				Quantity: item.Quantity,
				Price:    item.Product.Price,
				Product: dto.ProductResponse{
					ID:          item.Product.ID,
					CategoryID:  item.Product.CategoryID,
					Name:        item.Product.Name,
					Description: item.Product.Description,
					Price:       item.Product.Price,
					Stock:       item.Product.Stock,
					SKU:         item.Product.SKU,
					Category: dto.CategoryResponse{
						ID:          item.Product.Category.ID,
						Name:        item.Product.Category.Name,
						Description: item.Product.Category.Description,
						IsActive:    item.Product.Category.IsActive,
					},
				},
			}
		}

		if err := tx.Where("cart_id = ?", cart.ID).Delete(&models.CartItem{}).Error; err != nil {
			return errors.New("Failed to clear cart items")
		}

		orderResponse = &dto.OrderResponse{
			ID:          order.ID,
			UserID:      order.UserID,
			Status:      string(order.Status),
			TotalAmount: order.TotalAmount,
			OrderItems:  orderItemResponses,
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	if os.publisher != nil {
		metadata := map[string]string{
			"user_id":  strconv.Itoa(userID),
			"order_id": strconv.Itoa(orderResponse.ID),
		}
		if pubErr := os.publisher.Publish("order.created", orderResponse, metadata); pubErr != nil {
			log.Warn().Err(pubErr).Int("order_id", orderResponse.ID).Msg("Failed to publish order.created event to SQS")
		} else {
			log.Info().Int("order_id", orderResponse.ID).Msg("Successfully published order.created event to SQS")
		}
	}

	return orderResponse, nil
}

func (os *OrderService) GetOrders(userID int) ([]dto.OrderResponse, error) {
	var orders []models.Order
	if err := os.db.Preload("OrderItems.Product.Category").Where("user_id = ?", userID).Order("created_at desc").Find(&orders).Error; err != nil {
		return nil, err
	}

	result := make([]dto.OrderResponse, len(orders))
	for i := range orders {
		result[i] = *os.convertToOrderResponse(&orders[i])
	}
	return result, nil
}

func (os *OrderService) GetOrderByID(userID int, orderID int) (*dto.OrderResponse, error) {
	var order models.Order
	if err := os.db.Preload("OrderItems.Product.Category").Where("id = ? AND user_id = ?", orderID, userID).First(&order).Error; err != nil {
		return nil, errors.New("Order Not Found")
	}
	return os.convertToOrderResponse(&order), nil
}

func (os *OrderService) convertToOrderResponse(order *models.Order) *dto.OrderResponse {
	items := make([]dto.OrderItemResponse, len(order.OrderItems))
	for i, item := range order.OrderItems {
		items[i] = dto.OrderItemResponse{
			ID:       item.ID,
			Quantity: item.Quantity,
			Price:    item.Price,
			Product: dto.ProductResponse{
				ID:          item.Product.ID,
				CategoryID:  item.Product.CategoryID,
				Name:        item.Product.Name,
				Description: item.Product.Description,
				Price:       item.Product.Price,
				Stock:       item.Product.Stock,
				SKU:         item.Product.SKU,
				Category: dto.CategoryResponse{
					ID:          item.Product.Category.ID,
					Name:        item.Product.Category.Name,
					Description: item.Product.Category.Description,
					IsActive:    item.Product.Category.IsActive,
				},
			},
		}
	}

	return &dto.OrderResponse{
		ID:          order.ID,
		UserID:      order.UserID,
		Status:      string(order.Status),
		TotalAmount: order.TotalAmount,
		OrderItems:  items,
	}
}
