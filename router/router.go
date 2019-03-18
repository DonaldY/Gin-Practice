package router

import (
	"Gin-Practice/middleware"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {

	r := gin.New()

	r.Use(middleware.CORSMiddleware())

	return r
}
