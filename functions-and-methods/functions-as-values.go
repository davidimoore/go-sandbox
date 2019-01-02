package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(6, 3)
}

func main() {
	sumSquareRoot := func(x, y float64) float64 {
		return math.Sqrt(x + y)
	}
	fmt.Println(sumSquareRoot(20, 5))

	fmt.Println(compute(sumSquareRoot))
}
