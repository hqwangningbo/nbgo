package main

import (
	"fmt"
	"github.com/hqwangningbo/nbgo/cmd"
	"github.com/hqwangningbo/nbgo/utils"
)

// @title nbgo学习go实战
// @version v0.0.1
// @description 还在努力学习go ...ing
func main() {
	defer cmd.Clean()
	cmd.Start()

	token, _ := utils.GenerateToken(1, "zs")
	fmt.Println(token)
}
