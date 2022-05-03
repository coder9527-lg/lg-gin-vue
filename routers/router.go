package routers

import (
	"github/coder9527-lg/lg-gin-vue/controller"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	r.GET("/api/user/register", controller.Register)
	r.Run()
}
