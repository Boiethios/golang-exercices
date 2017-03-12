/// From https://tour.golang.org/flowcontrol/8

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	ans := x / 2
	previous := .0
	same := func(x float64, y float64) bool {
		const precision = 1000000
		return math.Ceil(x * precision) == math.Ceil(y * precision)
	}
	
	for ! same(ans, previous) {
		previous = ans
		ans = ans - ((ans * ans - x) / (2 * ans))
	}
	return ans
}

func main() {
	fmt.Println("my Sqrt:   ", Sqrt(2))
	fmt.Println("math.Sqrt: ", math.Sqrt(2))
}

