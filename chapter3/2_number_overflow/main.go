package main

import (
	"fmt"
	"math"
)

func main() {
	var counter int32 = math.MaxInt32
	fmt.Printf("counter: %d\n", counter)
	fmt.Printf("incremented counter: %d\n", Inc32(counter))
}

func Inc32(counter int32) int32 {
	if counter == math.MaxInt32 {
		panic("int32 overflow!")
	}
	return counter + 1
}

func AddInt(a, b int) int {
	if (b > 0 && a > math.MaxInt-b) || (b < 0 && a < math.MinInt-b) {
		panic("int overflow")
	}
	return a + b
}
