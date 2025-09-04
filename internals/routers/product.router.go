package routers

import (
	"github.com/federus1105/daysatu/internals/handlers"
	"github.com/federus1105/daysatu/internals/repositories"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitProductRouter(router *gin.Engine, db *pgxpool.Pool) {
	productRouter := router.Group("/products")
	productRepository := repositories.NewProductRepository(db)
	productHandler := handlers.NewProductHandler(productRepository)

	productRouter.POST("/products", productHandler.AddNewProduct)
	productRouter.PATCH("/products/:id", productHandler.GetProductWithparam)
}
