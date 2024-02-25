package controllers

import (
	"codezas-pos/db"
	"codezas-pos/dto"
	"codezas-pos/entity"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Account struct{}

func (c Account) FindAll(ctx *gin.Context) {
	var accounts []entity.Account
	db.Connection.Find(&accounts)

	var result []dto.AccountResponse
	for _, acc := range accounts {
		result = append(result, dto.AccountResponse{
			ID:        acc.ID,
			AccountId: acc.AccountId,
		})
	}

	ctx.JSON(http.StatusOK, result)
}
func (c Account) FindOne(ctx *gin.Context) {
	account_id := ctx.Param("id")
	print(account_id)
	var accounts entity.Account
	if err := db.Connection.First(&accounts, "account_id=? and status=?", account_id, 1).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, dto.AccountResponse{
		ID:        accounts.ID,
		AccountId: accounts.AccountId,
	})
}

func (c Account) Create(ctx *gin.Context) {
	var form dto.AccountRequest
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accounts := entity.Account{
		AccountId: form.AccountId,
		Status:    form.Status,
	}

	if err := db.Connection.Create(&accounts).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.AccountResponse{
		ID:        accounts.ID,
		AccountId: accounts.AccountId,
	})
}
