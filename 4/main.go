package main

import (
    "fmt"
       )

func IsPrime(number int64) bool {
    for x := int64(2); x < number; x++ {
        if number % x == 0 {
            return false
        }
    }
    return true
}

func main() {
    count := 0
    sum := int64(0)
    for x := int64(2); count < 1000; x++ {
        if IsPrime(x) {
            count++
            sum += x
        }
    }
    fmt.Println(sum)
}
