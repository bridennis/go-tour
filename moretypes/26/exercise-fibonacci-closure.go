/*
	Упражнение: Fibonacci closure
		https://tour.golang.org/moretypes/26
*/

package main

import "fmt"

func fibonacci() func() int {
	firstVal := -1
	secondVal := 1
	return func() int {
		secondVal = firstVal + secondVal
		firstVal = secondVal - firstVal
		return secondVal
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
