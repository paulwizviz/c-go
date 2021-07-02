package mylib_test

import (
	"fmt"

	"github.com/paulwizviz/c-go/mylib"
)

func ExampleMultiply() {
	result := mylib.Multiply(2, 3)
	fmt.Println(result)
	// Output:
	// 6
}

func ExampleSum() {
	result := mylib.Sum(1, 1)
	fmt.Println(result)
	// Output:
	// 2
}
