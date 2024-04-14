package main

import (
	"fmt"
	"math"
)

func main() {
	var counter int32 = math.MaxInt32
	fmt.Printf("counter: %d\n", counter)
	counter++
	fmt.Printf("incremented counter: %d\n", counter)
}
