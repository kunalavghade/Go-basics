package server

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kunalavghade/Go-basics/go-api/internal/utils"
)

func (s *Server) authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get token from Authorization header
		authorization := c.GetHeader("Authorization")
		if authorization == "" {
			utils.UnauthorizedResponse(c, "Authorization header is required", nil)
			c.Abort()
			return
		}

		// Split the token from the header
		tokenParts := strings.Split(authorization, " ")

		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			utils.UnauthorizedResponse(c, "Authorization header is required", nil)
			c.Abort()
			return
		}

		// Check if token is valid
		claims, err := utils.ValidateToken(tokenParts[1], s.config.JWT.Secret)
		if err != nil {
			utils.UnauthorizedResponse(c, "Invalid token", nil)
			c.Abort()
			return
		}

		// Set user in context
		c.Set("user_id", claims.UserID)
		c.Set("user_role", claims.Role)
		c.Set("user_email", claims.Email)

		c.Next()
	}
}

func (s *Server) adminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("user_role")
		if !exists {
			utils.UnauthorizedResponse(c, "Authorization header is required", nil)
			c.Abort()
			return
		}
		if userRole != "admin" {
			utils.ForbiddenResponse(c, "user is not admin", nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
