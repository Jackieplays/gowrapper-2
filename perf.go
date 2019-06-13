package main

import (
    "fmt"
    "time"
)

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void show() {

}

*/
// #cgo LDFLAGS: -lstdc++
import "C"

//import "fmt"

func show() {

}

func main() {
    now := time.Now()
    for i := 0; i < 100000000; i = i + 1 {
        C.show()
    }
    end_time := time.Now()

    var dur_time time.Duration = end_time.Sub(now)
   var elapsed_sec float64 = dur_time.Seconds()
 
    fmt.Printf("cgo show function elasped  \nelapsed %f seconds\n",
        elapsed_min, elapsed_sec, elapsed_nano)

    now = time.Now()
    for i := 0; i < 100000000; i = i + 1 {
        show()
    }
    end_time = time.Now()

    dur_time = end_time.Sub(now)
    
    elapsed_sec = dur_time.Seconds()
   
    fmt.Printf("go show function elasped  \nelapsed %f seconds \n",
        elapsed_min, elapsed_sec, elapsed_nano)

    var input string
    fmt.Scanln(&input)
}
