package mylib

// #include <stdio.h>
//
// int
// multiply(int lhs, int rhs) {
//    return lhs * rhs;
//}
import "C"

func Multiply(lhs int, rhs int) int {
	result := C.multiply(C.int(lhs), C.int(rhs))
	return int(result)
}
