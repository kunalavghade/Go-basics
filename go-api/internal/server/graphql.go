package server

import (
	"context"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/kunalavghade/Go-basics/go-api/graph"
	"github.com/kunalavghade/Go-basics/go-api/graph/resolver"
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

	schema := graph.NewExecutableSchema(
		graph.Config{
			Resolvers: rvr,
		},
	)
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

func (s *Server) GraphQLHandler() gin.HandlerFunc {
	h := s.createGQLHandler()
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		for k, v := range c.Keys {
			ctx = context.WithValue(ctx, k, v)
		}
		h.ServeHTTP(c.Writer, c.Request.WithContext(ctx))
	}
}

func (s *Server) PlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL Playground", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func (s *Server) GraphQLMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		userID, _ := c.Get("user_id")
		userEamil, _ := c.Get("user_email")
		userRole, _ := c.Get("user_role")

		ctx := context.WithValue(c.Request.Context(), "user_id", userID)
		ctx = context.WithValue(ctx, "user_email", userEamil)
		ctx = context.WithValue(ctx, "user_role", userRole)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
