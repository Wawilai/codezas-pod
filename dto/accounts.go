package dto

type AccountRequest struct {
	AccountId uint `form:"account_id" binding:"required"`
	Status    uint `form:"status" binding:"required"`
}

type AccountResponse struct {
	ID        uint `json:"id"`
	AccountId uint `json:"account_id"`
}
