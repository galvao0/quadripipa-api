package router

import "github.com/gin-gonic/gin"

func Init() {
	r := gin.Default()

	InitRoutes(r)

	r.Run(":8080")
}
