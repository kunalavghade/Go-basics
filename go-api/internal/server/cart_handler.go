package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kunalavghade/Go-basics/go-api/internal/dto"
	"github.com/kunalavghade/Go-basics/go-api/internal/utils"
)

func (s *Server) GetCartHandler(c *gin.Context) {
	userId := c.MustGet("user_id").(int)
	cart, err := s.cartService.GetCart(userId)
	if err != nil {
		utils.NotFoundResponse(c, "Cart Not Found", err)
		return
	}
	utils.SucessResponse(c, "Cart retrieved successfully", cart)
}

func (s *Server) AddToCartHandler(c *gin.Context) {
	var req dto.AddToCartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request data", err)
		return
	}
	userId := c.MustGet("user_id").(int)
	cart, err := s.cartService.AddToCart(userId, &req)
	if err != nil {
		utils.BadRequestResponse(c, err.Error(), err)
		return
	}
	utils.SucessResponse(c, "Item added to cart successfully", cart)
}

func (s *Server) UpdateCartHandler(c *gin.Context) {
	var req dto.UpdateCartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request data", err)
		return
	}
	userId := c.MustGet("user_id").(int)
	cart, err := s.cartService.UpdateCart(userId, &req)
	if err != nil {
		utils.BadRequestResponse(c, err.Error(), err)
		return
	}
	utils.SucessResponse(c, "Cart updated successfully", cart)
}

func (s *Server) RemoveFromCartHandler(c *gin.Context) {
	var req dto.RemoveFromCartRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request data", err)
		return
	}
	userId := c.MustGet("user_id").(int)
	cart, err := s.cartService.RemoveFromCart(userId, &req)
	if err != nil {
		utils.BadRequestResponse(c, err.Error(), err)
		return
	}
	utils.SucessResponse(c, "Item removed from cart successfully", cart)
}
