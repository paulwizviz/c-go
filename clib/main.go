package main

// #cgo LDFLAGS: -L . -lsum
// #include "sum.h"
import "C"

func main() {
	C.sum(1, 2)
}
