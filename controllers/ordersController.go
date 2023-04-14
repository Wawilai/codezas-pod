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

type Order struct{}

func (o Order) FindAll(ctx *gin.Context) {
	var orders []entity.Order
	db.Connection.Preload("Products").Find(&orders)

	var result []dto.OrderResponse
	for _, order := range orders {
		orderResult := dto.OrderResponse{
			ID:    order.ID,
			Name:  order.Name,
			Tel:   order.Tel,
			Email: order.Email,
		}
		var products []dto.OrderProductResponse
		for _, product := range order.Products {
			products = append(products, dto.OrderProductResponse{
				ID:       product.ID,
				SKU:      product.SKU,
				Name:     product.Name,
				Price:    product.Price,
				Quantity: product.Quantity,
				Image:    product.Image,
			})
		}
		orderResult.Products = products
		result = append(result, orderResult)
	}

	ctx.JSON(http.StatusOK, result)
}

func (o Order) FindOne(ctx *gin.Context) {
	id := ctx.Param("id")
	var order entity.Order

	query := db.Connection.Preload("Products").First(&order, id)
	if err := query.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	result := dto.OrderResponse{
		ID:    order.ID,
		Name:  order.Name,
		Tel:   order.Tel,
		Email: order.Email,
	}
	var products []dto.OrderProductResponse
	for _, product := range order.Products {
		products = append(products, dto.OrderProductResponse{
			ID:       product.ID,
			SKU:      product.SKU,
			Name:     product.Name,
			Price:    product.Price,
			Quantity: product.Quantity,
			Image:    product.Image,
		})
	}
	result.Products = products

	ctx.JSON(http.StatusOK, result)
}

func (o Order) Create(ctx *gin.Context) {
	var form dto.OrderRequest
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var order entity.Order
	var orderItems []entity.OrderItem
	for _, product := range form.Products {
		orderItems = append(orderItems, entity.OrderItem{
			SKU:      product.SKU,
			Name:     product.Name,
			Price:    product.Price,
			Quantity: product.Quantity,
			Image:    product.Image,
		})
	}
	order.Name = form.Name
	order.Tel = form.Tel
	order.Email = form.Email
	order.Products = orderItems
	db.Connection.Create(&order)

	result := dto.OrderResponse{
		ID:    order.ID,
		Name:  order.Name,
		Tel:   order.Tel,
		Email: order.Email,
	}
	var products []dto.OrderProductResponse
	for _, product := range order.Products {
		products = append(products, dto.OrderProductResponse{
			ID:       product.ID,
			SKU:      product.SKU,
			Name:     product.Name,
			Price:    product.Price,
			Quantity: product.Quantity,
			Image:    product.Image,
		})
	}
	result.Products = products

	ctx.JSON(http.StatusCreated, result)
}
