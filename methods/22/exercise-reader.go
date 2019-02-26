/*
	Упражнение: Reader
		https://go-tour-ru-ru.appspot.com/methods/22
*/

package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = "A"[0]
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
