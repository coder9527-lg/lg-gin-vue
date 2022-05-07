package main

import (
	"fmt"
	"os"

	"github.com/coder9527-lg/lg-gin-vue/cmd"
)

func test(s int) {
	if s/5 == 0 {
		fmt.Println("yes")
	}
}

func main() {
	if err := cmd.Execute(); err != nil {
		println("start fail:", err.Error())
	}
	os.Exit(-1)
}
