package main

// #include <stdio.h>
//
// void sum(int lhs, int rhs) {
//    printf("sum of %d and %d is %d.\n", lhs, rhs, lhs+rhs);
//}
import "C"

func main() {
	C.sum(1, 2)
}
