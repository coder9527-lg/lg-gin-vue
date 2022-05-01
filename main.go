package main

import (
	"github/coder9527-lg/lg-gin-vue/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("ping", controller.Ping)
	r.GET("/api/auth/register", controller.Register)
	r.Run()
}
