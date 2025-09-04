package routers

import (
	"github.com/federus1105/daysatu/internals/handlers"
	"github.com/gin-gonic/gin"
)

func InitPingRouter(router *gin.Engine) {
	pingRouter := router.Group("/ping")
	ph := handlers.NewPingHandler()

	pingRouter.GET("", ph.GetPing)
	pingRouter.GET("/:id/:param2", ph.GetPingWithParam)
	pingRouter.POST("", ph.PostPing)
}
