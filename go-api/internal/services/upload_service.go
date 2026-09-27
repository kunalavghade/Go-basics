package services

import (
	"errors"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/kunalavghade/Go-basics/go-api/internal/interfaces"
)

type UploadService struct {
	provider interfaces.UploadProvider
}

func NewUploadService(provider interfaces.UploadProvider) *UploadService {
	return &UploadService{provider: provider}
}

func (s *UploadService) UploadProductImage(productId int, file *multipart.FileHeader) (string, error) {
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !isValidImageExt(ext) {
		return "", errors.New("invalid file type")
	}
	if file.Size > 5*1024*1024 {
		return "", errors.New("file size exceeds limit")
	}
	path := fmt.Sprintf("products/%d/%s", productId, file.Filename)

	if _, err := s.provider.UploadFile(file, path); err != nil {
		return "", err
	}
	return path, nil
}

func isValidImageExt(ext string) bool {
	return ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif" || ext == ".webp" || ext == ".bmp"
}
