package main

// #cgo LDFLAGS: -L../clib -lsum
// #cgo CFLAGS: -I../clib
// #include "sum.h"
import "C"

func main() {
	C.sum(1, 1)
}
