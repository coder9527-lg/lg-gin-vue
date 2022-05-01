package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	//get the parameter
	name := c.PostForm("name")
	password := c.PostForm("password")
	phone := c.PostForm("phone")

	//Data verification
	//Create user
	//Return result
	c.JSON(http.StatusOK, gin.H{
		"msg": "Register success",
	})
}
