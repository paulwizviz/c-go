package mylib

// #cgo LDFLAGS: -L./ -lsum
// #cgo CFLAGS: -I./
// #include "sum.h"
import "C"

func Sum(lhs int, rhs int) int {
	result := C.sum(C.int(lhs), C.int(rhs))
	return int(result)
}
