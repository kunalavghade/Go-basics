package server

import (
	"context"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/kunalavghade/Go-basics/go-api/graph"
	"github.com/kunalavghade/Go-basics/go-api/graph/resolver"
	"github.com/kunalavghade/Go-basics/go-api/internal/utils"
	"github.com/vektah/gqlparser/v2/ast"
)

func (s *Server) createGQLHandler() *handler.Server {
	rvr := resolver.NewResolver(
		s.authService,
		s.userService,
		s.productService,
		s.cartService,
		s.orderService,
	)

	cfg := graph.Config{
		Resolvers: rvr,
	}

	// @auth directive handler: checks if user is authenticated for protected fields
	cfg.Directives.Auth = func(ctx context.Context, obj any, next graphql.Resolver) (res any, err error) {
		if _, err := resolver.GetUserIDByContext(ctx); err != nil {
			return nil, resolver.ErrUnauthorized
		}
		return next(ctx)
	}

	schema := graph.NewExecutableSchema(cfg)
	srv := handler.New(schema)
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})
	srv.SetQueryCache(lru.New[*ast.QueryDocument](1024))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})
	return srv
}

// GraphQLAuthMiddleware optionally extracts JWT authentication claims without blocking unauthenticated requests
func (s *Server) GraphQLAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		if authorization != "" {
			tokenParts := strings.Split(authorization, " ")
			if len(tokenParts) == 2 && tokenParts[0] == "Bearer" {
				claims, err := utils.ValidateToken(tokenParts[1], s.config.JWT.Secret)
				if err == nil {
					c.Set("user_id", claims.UserID)
					c.Set("user_role", claims.Role)
					c.Set("user_email", claims.Email)
				}
			}
		}

		ctx := c.Request.Context()
		if userID, exists := c.Get("user_id"); exists {
			ctx = context.WithValue(ctx, "user_id", userID)
		}
		if userRole, exists := c.Get("user_role"); exists {
			ctx = context.WithValue(ctx, "user_role", userRole)
		}
		if userEmail, exists := c.Get("user_email"); exists {
			ctx = context.WithValue(ctx, "user_email", userEmail)
		}
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

func (s *Server) GraphQLHandler() gin.HandlerFunc {
	h := s.createGQLHandler()
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func (s *Server) PlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL Playground", "/graphql")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
