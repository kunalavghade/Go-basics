package resolver

import "github.com/kunalavghade/Go-basics/go-api/internal/services"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.

type Resolver struct {
	authService    *services.AuthService
	userService    *services.UserService
	productService *services.ProductService
	cartService    *services.CartService
	orderService   *services.OrderService
}

func NewResolver(
	authService *services.AuthService,
	userService *services.UserService,
	productService *services.ProductService,
	cartService *services.CartService,
	orderService *services.OrderService,
) *Resolver {
	return &Resolver{
		authService:    authService,
		userService:    userService,
		productService: productService,
		cartService:    cartService,
		orderService:   orderService,
	}
}
