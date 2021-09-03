/*
	Упражнение: срезы
		https://tour.golang.org/moretypes/18
*/

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	area := make([][]uint8, dx)
	for i := range area {
		area[i] = make([]uint8, dy)
		for j := range area[i] {
			area[i][j] = uint8(i * j)
		}
	}
	return area
}

func main() {
	pic.Show(Pic)
}
