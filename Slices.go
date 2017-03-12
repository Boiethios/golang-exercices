package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	ans := make([][]uint8, dy)
	
	for y := range ans {
		ans[y] = make([]uint8, dx)
		for x := range ans[y] {
			ans[y][x] = uint8(x * y)
		}
	}
	return ans
}

func main() {
	pic.Show(Pic)
}

