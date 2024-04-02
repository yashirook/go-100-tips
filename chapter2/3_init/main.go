package main

import (
	"fmt"
	_ "init/pkg/foo"
	"init/pkg/redis"
)

var a = func() int {
	fmt.Println("var")
	return 0
}()

func init() {
	fmt.Println("init first")
}

func init() {
	fmt.Println("init second")
}

func main() {
	fmt.Println("main")
	redis.Store("test", "testvalue")
}
