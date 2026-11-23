package router

import (
	"github.com/galvao0/quadripipa-api/handler"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/greeting", handler.Greeting)
	}
}
