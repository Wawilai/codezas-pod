package dto

type CategoryRequest struct {
	Name string `form:"name" binding:"required"`
}

type CategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
