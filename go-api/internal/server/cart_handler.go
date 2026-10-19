package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kunalavghade/Go-basics/go-api/internal/dto"
	"github.com/kunalavghade/Go-basics/go-api/internal/utils"
)

// GetCartHandler godoc
// @Summary Get user cart
// @Description Retrieve cart contents and total for the authenticated user
// @Tags Cart
// @Produce json
// @Security BearerAuth
// @Success 200 {object} utils.Response{data=dto.CartResponse}
// @Failure 404 {object} utils.Response
// @Router /api/cart/ [get]
func (s *Server) GetCartHandler(c *gin.Context) {
	userId := c.MustGet("user_id").(int)
	cart, err := s.cartService.GetCart(userId)
	if err != nil {
		utils.NotFoundResponse(c, "Cart Not Found", err)
		return
	}
	utils.SucessResponse(c, "Cart retrieved successfully", cart)
}

// AddToCartHandler godoc
// @Summary Add item to cart
// @Description Add product to cart or increase quantity
// @Tags Cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.AddToCartRequest true "Cart Item Info"
// @Success 200 {object} utils.Response{data=dto.CartResponse}
// @Failure 400 {object} utils.Response
// @Router /api/cart/add [post]
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

// UpdateCartHandler godoc
// @Summary Update cart item
// @Description Update quantity of item in cart (quantity <= 0 deletes item)
// @Tags Cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.UpdateCartRequest true "Update Cart Item Info"
// @Success 200 {object} utils.Response{data=dto.CartResponse}
// @Failure 400 {object} utils.Response
// @Router /api/cart/update [put]
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

// RemoveFromCartHandler godoc
// @Summary Remove item from cart
// @Description Remove a product from the user cart
// @Tags Cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param product_id query int false "Product ID (query)"
// @Param request body dto.RemoveFromCartRequest false "Product ID (JSON body)"
// @Success 200 {object} utils.Response{data=dto.CartResponse}
// @Failure 400 {object} utils.Response
// @Router /api/cart/remove [delete]
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
