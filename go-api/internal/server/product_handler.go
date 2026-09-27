package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kunalavghade/Go-basics/go-api/internal/dto"
	"github.com/kunalavghade/Go-basics/go-api/internal/services"
)

func (s *Server) CreateProductHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.CreateProductRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		productService := services.NewProductService(s.DB)
		product, err := productService.CreateProduct(&req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, product)
	}
}

func (s *Server) GetProductsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
		productService := services.NewProductService(s.DB)
		products, paginationMeta := productService.GetProducts(page, limit)
		if products == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": products, "meta": paginationMeta})
	}
}

func (s *Server) GetProductByIDHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		productService := services.NewProductService(s.DB)
		product, err := productService.GetProduct(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, product)
	}
}

func (s *Server) UpdateProductHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var req dto.UpdateProductRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		productService := services.NewProductService(s.DB)
		product, err := productService.UpdateProduct(id, &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, product)
	}
}

func (s *Server) DeleteProductHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		productService := services.NewProductService(s.DB)
		if err := productService.DeleteProduct(id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusNoContent)
	}
}

func (s *Server) CreateCategoryHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.CreateCategoryRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		productService := services.NewProductService(s.DB)
		category, err := productService.CreateCategory(&req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, category)
	}
}

func (s *Server) GetCategoriesHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		productService := services.NewProductService(s.DB)
		categories, err := productService.GetGategories()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, categories)
	}
}

func (s *Server) UpdateCategoryHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var req dto.UpdateCategoryRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		productService := services.NewProductService(s.DB)
		category, err := productService.UpdateCategory(id, &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, category)
	}
}

func (s *Server) DeleteCategoryHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		productService := services.NewProductService(s.DB)
		if err := productService.DeleteCategory(id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusNoContent)
	}
}
