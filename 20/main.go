package main

import (
    "fmt"
    "os"
    "strings"
    "bufio"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func LowerCaseLine(line string) string {
    return strings.ToLower(line)
}


func main() {
    file, err := os.Open(os.Args[1])
    check(err)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(LowerCaseLine(scanner.Text()))
    }
}
