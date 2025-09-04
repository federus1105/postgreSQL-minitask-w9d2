package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/federus1105/daysatu/internals/repositories"
	"github.com/gin-gonic/gin"
)

type BooksHandler struct {
	sr *repositories.BooksRepository
}

func NewStudentHandler(sr *repositories.BooksRepository) *BooksHandler {
	return &BooksHandler{
		sr: sr,
	}
}
func (b *BooksHandler) GetBook(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		page = 1
	}
	limit := 5
	offset := (page - 1) * limit

	books, err := b.sr.GetBooksData(ctx.Request.Context(), offset, limit)
	if err != nil {
		log.Println("Internal server Error: ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"succes": false,
			"data":   []any{},
		})
		return
	}
	if len(books) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"succes": false,
			"data":   []any{},
			"page":   page,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"succes": true,
		"data":   books,
	})
}
