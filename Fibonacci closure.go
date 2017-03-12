package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	current := 0
	previous := 1
	
	return func() int {
		ans := current
		current = current + previous
		previous = ans
		return ans
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
