package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/kunalavghade/Go-basics/go-api/internal/dto"
	"github.com/kunalavghade/Go-basics/go-api/internal/utils"
)

// CreateOrderHandler godoc
// @Summary Checkout & create order
// @Description Create an order from the user's cart, deduct stock, and empty cart
// @Tags Orders
// @Produce json
// @Security BearerAuth
// @Success 201 {object} utils.Response{data=dto.OrderResponse}
// @Failure 400 {object} utils.Response
// @Router /api/orders/ [post]
func (s *Server) CreateOrderHandler(c *gin.Context) {
	userId := c.MustGet("user_id").(int)
	order, err := s.orderService.CreateOrder(userId)
	if err != nil {
		utils.BadRequestResponse(c, err.Error(), err)
		return
	}
	utils.CreatedResponse(c, "Order created successfully", order)
}

// GetOrdersHandler godoc
// @Summary List user orders
// @Description Fetch all orders placed by the authenticated user
// @Tags Orders
// @Produce json
// @Security BearerAuth
// @Success 200 {object} utils.Response{data=[]dto.OrderResponse}
// @Failure 500 {object} utils.Response
// @Router /api/orders/ [get]
func (s *Server) GetOrdersHandler(c *gin.Context) {
	userId := c.MustGet("user_id").(int)
	orders, err := s.orderService.GetOrders(userId)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to fetch orders", err)
		return
	}
	utils.SucessResponse(c, "Orders fetched successfully", orders)
}

// GetOrderByIDHandler godoc
// @Summary Get order by ID
// @Description Fetch order details including ordered items and product info
// @Tags Orders
// @Produce json
// @Security BearerAuth
// @Param id path int true "Order ID"
// @Success 200 {object} utils.Response{data=dto.OrderResponse}
// @Failure 400 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Router /api/orders/{id} [get]
func (s *Server) GetOrderByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.BadRequestResponse(c, "Invalid order ID", err)
		return
	}

	userId := c.MustGet("user_id").(int)
	order, err := s.orderService.GetOrderByID(userId, id)
	if err != nil {
		utils.NotFoundResponse(c, "Order Not Found", err)
		return
	}
	utils.SucessResponse(c, "Order retrieved successfully", order)
}
