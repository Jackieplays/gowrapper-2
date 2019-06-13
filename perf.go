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
 
    fmt.Printf("cgo show function elapsed  \nelapsed %f seconds\n",
        elapsed_sec)

    now = time.Now()
    for i := 0; i < 100000000; i = i + 1 {
        show()
    }
    end_time = time.Now()

    dur_time = end_time.Sub(now)
    
    elapsed_sec = dur_time.Seconds()
   
    fmt.Printf("go show function elapsed  \nelapsed %f seconds \n",
        elapsed_sec)

    var input string
    fmt.Scanln(&input)
}

Output: 
cgo show function elapsed  
elapsed 6.610593 seconds
go show function elapsed  
elapsed 0.033000 seconds 

