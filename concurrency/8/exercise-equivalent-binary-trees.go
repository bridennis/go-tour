/*
	Упражнение: равнозначные двоичные деревья
		https://go-tour-ru-ru.appspot.com/concurrency/8
*/

package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var f func(*tree.Tree, chan int)
	f = func(t *tree.Tree, ch chan int) {
		if t.Left != nil {
			f(t.Left, ch)
		}
		ch <- t.Value
		if t.Right != nil {
			f(t.Right, ch)
		}
	}
	f(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := range ch1 {
		if i != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
