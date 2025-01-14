package dto

type CategoryRequest struct {
	ShopId string `json:"shop_id" binding:"required"`
	Name   string `json:"name" binding:"required,min=3,max=30"`
}

type UpdateCategoryRequest struct {
	Name string `json:"name" binding:"omitempty,min=3,max=30"`
}
