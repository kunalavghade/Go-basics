package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type PaginatedResponse struct {
	Response
	Meta PaginationMeta `json:"meta"`
}

type PaginationMeta struct {
	Page         int `json:"page"`
	Limit        int `json:"limit"`
	TotalPages   int `json:"total_pages"`
	TotalRecords int `json:"total_records"`
}

func SucessResponse(c *gin.Context, msg string, data interface{}) Response {
	return c.JSON(http.StatusOK, Response{
		Success: true,
		Message: msg,
		Data:    data,
	})
}

func CreatedResponse(c *gin.Context, msg string, data interface{}) Response {
	return c.JSON(http.StatusCreated, Response{
		Success: true,
		Message: msg,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, statusCode int, msg string, err error) Response {
	response := Response{
		Success: false,
		Message: msg,
	}
	if err != nil {
		response.Error = err.Error()
	}
	return c.JSON(statusCode, response)
}

func BadRequestResponse(c *gin.Context, msg string, err error) Response {
	return ErrorResponse(c, http.StatusBadRequest, msg, err)
}

func NotFoundResponse(c *gin.Context, msg string, err error) Response {
	return ErrorResponse(c, http.StatusNotFound, msg, err)
}

func UnprocessableEntityResponse(c *gin.Context, msg string, err error) Response {
	return ErrorResponse(c, http.StatusUnprocessableEntity, msg, err)
}

func UnauthorizedResponse(c *gin.Context, msg string, err error) Response {
	return ErrorResponse(c, http.StatusUnauthorized, msg, err)
}

func ForbiddenResponse(c *gin.Context, msg string, err error) Response {
	return ErrorResponse(c, http.StatusForbidden, msg, err)
}

func InternalServerErrorResponse(c *gin.Context, msg string, err error) Response {
	return ErrorResponse(c, http.StatusInternalServerError, msg, err)
}

func ConflictResponse(c *gin.Context, msg string, err error) Response {
	return ErrorResponse(c, http.StatusConflict, msg, err)
}

func PaginationResponse(c *gin.Context, data interface{}, page, limit, totalItems int) Response {
	totalPages := (totalItems + limit - 1) / limit

	response := Response{
		Success: true,
		Message: "Data retrieved successfully",
		Data: PaginatedResponse{
			Response: Response{
				Success: true,
				Message: "Data retrieved successfully",
				Data:    data,
			},
			Meta: PaginationMeta{
				Page:         page,
				Limit:        limit,
				TotalPages:   totalPages,
				TotalRecords: totalItems,
			},
		},
	}
	return c.JSON(http.StatusOK, response)
}
