package main

import (
	"fmt"
	"github.com/hqwangningbo/gogofly/cmd"
	"github.com/hqwangningbo/gogofly/utils"
)

// @title gogofly学习go实战
// @version v0.0.1
// @description 还在努力学习go ...ing
func main() {
	defer cmd.Clean()
	cmd.Start()

	token, _ := utils.GenerateToken(1, "zs")
	fmt.Println(token)
}
