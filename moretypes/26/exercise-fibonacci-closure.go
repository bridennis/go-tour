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
		nextVal := firstVal + secondVal
		firstVal = secondVal
		secondVal = nextVal
		return nextVal
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
