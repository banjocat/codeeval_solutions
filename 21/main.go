package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)

func SumLine (line string) int64 {
    var sum int64
    for _,c := range line {
        num, err := strconv.ParseInt(string(c), 10, 64)
        check(err)
        sum += num
    }
    return sum
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}



func main() {
    file, err := os.Open(os.Args[1])
    check(err)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(SumLine(scanner.Text()))
    }
}
