package utils

import (
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Pagination struct {
	Page       int         `json:"page"`
	PageSize   int         `json:"pageSize"`
	TotalItems int64       `json:"totalItems"`
	TotalPages int         `json:"totalPages"`
	Items      interface{} `json:"items"`
}

func PaginationParams(c *fiber.Ctx) (page, pageSize int) {
	pageStr := c.Query("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSizeStr := c.Query("pageSize", "10")
	pageSize, err = strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	return page, pageSize
}

func CreatePagination(page, pageSize int, totalItems int64, items interface{}) Pagination {
	totalPages := int(math.Ceil(float64(totalItems) / float64(pageSize)))

	return Pagination{
		Page:       page,
		PageSize:   pageSize,
		TotalItems: totalItems,
		TotalPages: totalPages,
		Items:      items,
	}
}
