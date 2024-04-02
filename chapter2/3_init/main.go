package main

import (
	"fmt"
	"init/pkg/redis"
)

var a = func() int {
	fmt.Println("var")
	return 0
}()

func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("main")
	redis.Store("test", "testvalue")
}
