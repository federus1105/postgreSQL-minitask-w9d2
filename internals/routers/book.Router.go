package routers

import (
	"github.com/federus1105/daysatu/internals/handlers"
	"github.com/federus1105/daysatu/internals/repositories"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitBooksRouter(router *gin.Engine, db *pgxpool.Pool) {
	studentRouter := router.Group("/students")
	sr := repositories.NewBooksRepository(db)
	sh := handlers.NewBooksHandler(sr)

	studentRouter.GET("", sh.GetBook)
}
