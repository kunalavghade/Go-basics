package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kunalavghade/Go-basics/go-api/internal/dto"
	"github.com/kunalavghade/Go-basics/go-api/internal/utils"
)

// CreateProductHandler godoc
// @Summary Create product
// @Description Create a new product with category, price, and stock
// @Tags Products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateProductRequest true "Product info"
// @Success 201 {object} utils.Response{data=dto.ProductResponse}
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /api/products/ [post]
func (s *Server) CreateProductHandler(c *gin.Context) {
	var req dto.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request data", err)
		return
	}
	product, err := s.productService.CreateProduct(&req)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to create product", err)
		return
	}
	utils.CreatedResponse(c, "Product created successfully", product)
}

// GetProductsHandler godoc
// @Summary List products
// @Description Get paginated list of products
// @Tags Products
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /api/products/ [get]
func (s *Server) GetProductsHandler(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	products, paginationMeta := s.productService.GetProducts(page, limit)
	if products == nil {
		utils.InternalServerErrorResponse(c, "Failed to fetch products", nil)
		return
	}
	utils.SucessResponse(c, "Products fetched successfully", gin.H{"products": products, "meta": paginationMeta})
}

// GetProductByIDHandler godoc
// @Summary Get product by ID
// @Description Fetch details of a single product
// @Tags Products
// @Produce json
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Success 200 {object} utils.Response{data=dto.ProductResponse}
// @Failure 404 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /api/products/{id} [get]
func (s *Server) GetProductByIDHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := s.productService.GetProduct(id)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to fetch product", err)
		return
	}
	utils.SucessResponse(c, "Product fetched successfully", product)
}

// UpdateProductHandler godoc
// @Summary Update product
// @Description Update an existing product
// @Tags Products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Param request body dto.UpdateProductRequest true "Update data"
// @Success 200 {object} utils.Response{data=dto.ProductResponse}
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /api/products/{id} [put]
func (s *Server) UpdateProductHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req dto.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request data", err)
		return
	}
	product, err := s.productService.UpdateProduct(id, &req)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to update product", err)
		return
	}
	utils.SucessResponse(c, "Product updated successfully", product)
}

// DeleteProductHandler godoc
// @Summary Delete product
// @Description Remove a product by ID
// @Tags Products
// @Produce json
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Success 200 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /api/products/{id} [delete]
func (s *Server) DeleteProductHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := s.productService.DeleteProduct(id); err != nil {
		utils.InternalServerErrorResponse(c, "Failed to delete product", err)
		return
	}
	utils.SucessResponse(c, "Product deleted successfully", nil)
}

// CreateCategoryHandler godoc
// @Summary Create category
// @Description Create a new product category
// @Tags Categories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateCategoryRequest true "Category info"
// @Success 201 {object} utils.Response{data=dto.CategoryResponse}
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /api/categories/ [post]
func (s *Server) CreateCategoryHandler(c *gin.Context) {
	var req dto.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request data", err)
		return
	}
	category, err := s.productService.CreateCategory(&req)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to create category", err)
		return
	}
	utils.CreatedResponse(c, "Category created successfully", category)
}

// GetCategoriesHandler godoc
// @Summary List categories
// @Description Get all categories
// @Tags Categories
// @Produce json
// @Security BearerAuth
// @Success 200 {object} utils.Response{data=[]dto.CategoryResponse}
// @Failure 500 {object} utils.Response
// @Router /api/categories/ [get]
func (s *Server) GetCategoriesHandler(c *gin.Context) {
	categories, err := s.productService.GetGategories()
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to fetch categories", err)
		return
	}
	utils.SucessResponse(c, "Categories fetched successfully", categories)
}

// UpdateCategoryHandler godoc
// @Summary Update category
// @Description Update an existing category by ID
// @Tags Categories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Category ID"
// @Param request body dto.UpdateCategoryRequest true "Category data"
// @Success 200 {object} utils.Response{data=dto.CategoryResponse}
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /api/categories/{id} [put]
func (s *Server) UpdateCategoryHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req dto.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request data", err)
		return
	}
	category, err := s.productService.UpdateCategory(id, &req)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to update category", err)
		return
	}
	utils.SucessResponse(c, "Category updated successfully", category)
}

// DeleteCategoryHandler godoc
// @Summary Delete category
// @Description Remove a category by ID
// @Tags Categories
// @Produce json
// @Security BearerAuth
// @Param id path int true "Category ID"
// @Success 200 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /api/categories/{id} [delete]
func (s *Server) DeleteCategoryHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := s.productService.DeleteCategory(id); err != nil {
		utils.InternalServerErrorResponse(c, "Failed to delete category", err)
		return
	}
	utils.SucessResponse(c, "Category deleted successfully", nil)
}

// UploadProductImageHandler godoc
// @Summary Upload product image
// @Description Upload an image for a specific product
// @Tags Products
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Param image formData file true "Image file"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /api/products/{id}/image [post]
func (s *Server) UploadProductImageHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.BadRequestResponse(c, "Invalide Product ID", err)
		return
	}
	file, err := c.FormFile("image")
	if err != nil {
		utils.BadRequestResponse(c, "Invalide file", err)
		return
	}
	url, err := s.uploadService.UploadProductImage(int(id), file)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to upload image", err)
		return
	}

	if err := s.productService.AddProductImage(int(id), url, file.Filename); err != nil {
		utils.InternalServerErrorResponse(c, "Failed to add product image", err)
		return
	}
	utils.SucessResponse(c, "Image uploaded successfully", url)
}
