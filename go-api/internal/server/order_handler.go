package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kunalavghade/Go-basics/go-api/internal/utils"
)

func (s *Server) CreateOrderHandler(c *gin.Context) {
	userId := c.MustGet("user_id").(int)
	order, err := s.orderService.CreateOrder(userId)
	if err != nil {
		utils.BadRequestResponse(c, err.Error(), err)
		return
	}
	utils.CreatedResponse(c, "Order created successfully", order)
}

func (s *Server) GetOrdersHandler(c *gin.Context) {
	userId := c.MustGet("user_id").(int)
	orders, err := s.orderService.GetOrders(userId)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to fetch orders", err)
		return
	}
	utils.SucessResponse(c, "Orders fetched successfully", orders)
}

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
