package main

import (
	"fmt"

	"github/coder9527-lg/lg-gin-vue/db"
	"github/coder9527-lg/lg-gin-vue/routers"
)

func test(s int) {
	if s/5 == 0 {
		fmt.Println("yes")
	}
}

func main() {
	db.Run()
	routers.Run()
}
