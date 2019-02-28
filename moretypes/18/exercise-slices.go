/*
	Упражнение: срезы
		https://go-tour-ru-ru.appspot.com/moretypes/18
*/

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	area := make([][]uint8, dx)
	for i := 0; i < dx; i++ {
		area[i] = make([]uint8, dy)
		for j := 0; j < dy; j++ {
			area[i][j] = uint8(i * j)
		}
	}
	return area
}

func main() {
	pic.Show(Pic)
}
