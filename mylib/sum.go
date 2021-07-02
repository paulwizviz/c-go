package mylib

// #cgo LDFLAGS: -L./clib -lsum
// #cgo CFLAGS: -I./clib
// #include "sum.h"
import "C"

func Sum(lhs int, rhs int) int {
	result := C.sum(C.int(lhs), C.int(rhs))
	return int(result)
}
