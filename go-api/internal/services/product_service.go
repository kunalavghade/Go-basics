package services

import (
	"github.com/kunalavghade/Go-basics/go-api/internal/dto"
	"github.com/kunalavghade/Go-basics/go-api/internal/models"
	"github.com/kunalavghade/Go-basics/go-api/internal/utils"
	"gorm.io/gorm"
)

type ProductService struct {
	db *gorm.DB
}

func NewProductService(db *gorm.DB) *ProductService {
	return &ProductService{db: db}
}

func (s *ProductService) CreateCategory(req *dto.CreateCategoryRequest) (*dto.CategoryResponse, error) {
	category := models.Category{
		Name:        req.Name,
		Description: req.Description,
	}
	if err := s.db.Create(&category).Error; err != nil {
		return nil, err
	}
	return &dto.CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		IsActive:    category.IsActive,
	}, nil
}

func (s *ProductService) GetGategories() ([]dto.CategoryResponse, error) {
	var categories []models.Category
	if err := s.db.Where("is_active = ?", true).Find(&categories).Error; err != nil {
		return nil, err
	}
	categoryResponses := make([]dto.CategoryResponse, len(categories))
	for i := range categories {
		category := &categories[i]
		categoryResponses[i] = dto.CategoryResponse{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
			IsActive:    category.IsActive,
		}
	}
	return categoryResponses, nil
}

func (s *ProductService) UpdateCategory(id int, req *dto.UpdateCategoryRequest) (*dto.CategoryResponse, error) {
	var category models.Category
	if err := s.db.Where("id = ?", id).First(&category).Error; err != nil {
		return nil, err
	}
	if req.Name != "" {
		category.Name = req.Name
	}
	if req.Description != "" {
		category.Description = req.Description
	}
	category.IsActive = req.IsActive

	if err := s.db.Save(&category).Error; err != nil {
		return nil, err
	}
	return &dto.CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		IsActive:    category.IsActive,
	}, nil
}

func (s *ProductService) DeleteCategory(id int) error {
	var category models.Category
	if err := s.db.Where("id = ?", id).First(&category).Error; err != nil {
		return err
	}
	return s.db.Delete(&category).Error
}

func (s *ProductService) CreateProduct(req *dto.CreateProductRequest) (*dto.ProductResponse, error) {
	product := models.Product{
		Name:        req.Name,
		SKU:         req.SKU,
		Stock:       req.Stock,
		Price:       req.Price,
		CategoryID:  req.CategoryID,
		Description: req.Description,
	}

	if err := s.db.Create(&product).Error; err != nil {
		return nil, err
	}
	return s.GetProduct(product.ID)
}

func (s *ProductService) GetProducts(page, limit int) ([]*dto.ProductResponse, *utils.PaginationMeta) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit
	var products []*models.Product
	if err := s.db.Limit(limit).Offset(offset).Find(&products).Error; err != nil {
		return nil, nil
	}
	productResponses := make([]*dto.ProductResponse, len(products))
	for i, product := range products {
		productResponses[i] = s.convertToProductResponse(product)
	}
	var totalItems int64
	s.db.Model(&models.Product{}).Count(&totalItems)
	paginationMeta := &utils.PaginationMeta{
		Page:         page,
		Limit:        limit,
		TotalPages:   int(totalItems / int64(limit)),
		TotalRecords: int(totalItems),
	}
	return productResponses, paginationMeta
}

func (s *ProductService) GetProduct(id int) (*dto.ProductResponse, error) {
	var product models.Product
	if err := s.db.Preload("Category").Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}
	return s.convertToProductResponse(&product), nil
}

func (s *ProductService) UpdateProduct(id int, req *dto.UpdateProductRequest) (*dto.ProductResponse, error) {
	var product models.Product
	if err := s.db.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}
	if req.Name != "" {
		product.Name = req.Name
	}
	if req.Description != "" {
		product.Description = req.Description
	}
	if req.Price != 0 {
		product.Price = req.Price
	}
	if req.Stock != 0 {
		product.Stock = req.Stock
	}
	if req.CategoryID != 0 {
		product.CategoryID = req.CategoryID
	}
	if req.SKU != "" {
		product.SKU = req.SKU
	}
	if err := s.db.Save(&product).Error; err != nil {
		return nil, err
	}
	return s.GetProduct(product.ID)
}

func (s *ProductService) DeleteProduct(id int) error {
	var product models.Product
	if err := s.db.Where("id = ?", id).First(&product).Error; err != nil {
		return err
	}
	return s.db.Delete(&product).Error
}

func (s *ProductService) convertToProductResponse(product *models.Product) *dto.ProductResponse {
	images := make([]dto.ProductImageResponse, len(product.Images))
	for i := range product.Images {
		image := &product.Images[i]
		images[i] = dto.ProductImageResponse{
			ID:        image.ID,
			ProductID: image.ProductID,
			ImageURL:  image.ImageURL,
		}
	}
	return &dto.ProductResponse{
		ID:          product.ID,
		CategoryID:  product.CategoryID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		SKU:         product.SKU,
		Images:      images,
		Category: dto.CategoryResponse{
			ID:          product.Category.ID,
			Name:        product.Category.Name,
			Description: product.Category.Description,
			IsActive:    product.Category.IsActive,
		},
	}
}
