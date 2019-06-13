package main

// #cgo CFLAGS: -g -Wall
// #include <stdlib.h>
// #include"test1.h"
import "C"
import (
	"fmt"
)

func main() {
	C.prime(11)
	fmt.Println("yes done")
}


Output : Enter a positive integer:
 11 is a prime number.
yes done
