package main

import "fmt"

func main() {
    sum := 0
    i := 1
    j := 1
    k := 1
    for k < 4000000 {
        k  = i + j
        i = j
        j = k
        
        if (k % 2 == 0) {
            sum += k
            }
    }
    fmt.Println(sum)
}
