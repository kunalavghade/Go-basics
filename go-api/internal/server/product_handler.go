package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kunalavghade/Go-basics/go-api/internal/dto"
	"github.com/kunalavghade/Go-basics/go-api/internal/providers"
	"github.com/kunalavghade/Go-basics/go-api/internal/services"
	"github.com/kunalavghade/Go-basics/go-api/internal/utils"
)

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

func (s *Server) GetProductByIDHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := s.productService.GetProduct(id)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to fetch product", err)
		return
	}
	utils.SucessResponse(c, "Product fetched successfully", product)
}

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

func (s *Server) DeleteProductHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := s.productService.DeleteProduct(id); err != nil {
		utils.InternalServerErrorResponse(c, "Failed to delete product", err)
		return
	}
	utils.SucessResponse(c, "Product deleted successfully", nil)
}

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

func (s *Server) GetCategoriesHandler(c *gin.Context) {
	categories, err := s.productService.GetGategories()
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to fetch categories", err)
		return
	}
	utils.SucessResponse(c, "Categories fetched successfully", categories)
}

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

func (s *Server) DeleteCategoryHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := s.productService.DeleteCategory(id); err != nil {
		utils.InternalServerErrorResponse(c, "Failed to delete category", err)
		return
	}
	utils.SucessResponse(c, "Category deleted successfully", nil)
}

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
	uploadProvider := providers.NewLocalUploadProvider(s.config.Upload.Path)
	uploadService := services.NewUploadService(uploadProvider)

	url, err := uploadService.UploadProductImage(int(id), file)
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
