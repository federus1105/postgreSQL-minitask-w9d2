package routers

import (
	"net/http"

	"github.com/federus1105/daysatu/internals/middlewares"
	"github.com/federus1105/daysatu/internals/models"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitRouter(db *pgxpool.Pool) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.MyLogger)
	router.Use(middlewares.CORSMiddleware)

	InitPingRouter(router)
	InitBooksRouter(router, db)
	InitProductRouter(router, db)

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, models.Response{
			Message: "Rute salah",
			Status:  "Rute tidak ditemukan",
		})
	})
	return router
}
