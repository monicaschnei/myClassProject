package router

import "github.com/gin-gonic/gin"

func Post(router *gin.RouterGroup, path string, handler func(c *gin.Context)) {
	router.POST(path, handler)
}

func Get(router *gin.RouterGroup, path string, handler func(c *gin.Context)) {
	router.GET(path, handler)
}

func Put(router *gin.RouterGroup, path string, handler func(c *gin.Context)) {
	router.PUT(path, handler)
}
