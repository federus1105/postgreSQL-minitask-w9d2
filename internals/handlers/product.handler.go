package handlers

import (
	"net/http"

	"github.com/federus1105/daysatu/internals/models"
	"github.com/federus1105/daysatu/internals/repositories"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	pr *repositories.ProductRepository
}

func NewProductHandler(pr *repositories.ProductRepository) *ProductHandler {
	return &ProductHandler{
		pr: pr,
	}
}
func (p *ProductHandler) AddNewProduct(ctx *gin.Context) {
	var body models.Product
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}
	newProduct, err := p.pr.AddNewProduct(ctx.Request.Context(), body)
		if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
		})
		return
	}
		ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    newProduct,
	})
}
