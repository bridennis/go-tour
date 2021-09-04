/*
	Упражнение: Fibonacci closure
		https://tour.golang.org/moretypes/26
*/

package main

import "fmt"

func fibonacci() func() int {
	firstVal, secondVal := -1, 1
	return func() int {
		firstVal, secondVal = secondVal, firstVal+secondVal
		return secondVal
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
