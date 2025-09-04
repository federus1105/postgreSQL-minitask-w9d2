package handlers

import (
	"log"
	"net/http"

	"github.com/federus1105/daysatu/internals/models"
	"github.com/federus1105/daysatu/internals/utils"
	"github.com/gin-gonic/gin"
)

type PingHandler struct{}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

func (p *PingHandler) Getping(ctx *gin.Context) {
	requestId := ctx.GetHeader("X-Request-ID")
	contentType := ctx.GetHeader("Content-Type")
	ctx.JSON(http.StatusOK, gin.H{
		"message":     "pong",
		"requestId":   requestId,
		"contentType": contentType,
	})
}

func (p *PingHandler) GetWithParam(ctx *gin.Context) {
	pingId := ctx.Param("id")
	param2 := ctx.Param("param2")
	query := ctx.Query("query")
	ctx.JSON(http.StatusOK, gin.H{
		"param":  pingId,
		"param2": param2,
		"query":  query,
	})
}

func (p *PingHandler) PostPing(ctx *gin.Context) {
	body := models.Body{}
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}
	if err := utils.ValidateBody(body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	log.Println(body)
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"body":    body,
	})
}
