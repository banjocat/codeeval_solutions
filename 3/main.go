package main

import (
    "fmt"
    "strconv"
       )

func IsPrime(number int) bool {
    for x := 2; x < number; x++ {
        if number % x == 0 {
            return false
        }
    }
    return true
}

func ReverseString  (s string) string {
    var reverse string
    for _, val := range s {
        reverse = string(val) + reverse;
    }
    return reverse
}


func IsPalindrome(number int) bool {
    check :=  strconv.FormatInt(int64(number), 10)
    reverse := ReverseString(check)
    return check == reverse
}

func main ()  {
    for x := 999; x >= 1; x-- {
        if IsPrime(x) && IsPalindrome(x) {
            fmt.Println(x)
            return
        }
    }
}
