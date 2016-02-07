package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)

type Lines []string

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func (list Lines) LongestLines (check string) Lines {
    for n,line := range list {
        if len(check) >=  len(line) {
            return append(list[:n], append(Lines{check}, list[n:]...)...)
        }
    }
    return append(list,check)
}

func (list Lines) ReduceLengthIfNeeded(n int64) Lines {
    if int64(len(list)) <= n {
        return list
    }
    return list[:len(list) -1]
}

func main() {
    file,err := os.Open(os.Args[1])
    check(err)
    scanner := bufio.NewScanner(file)
    scanner.Scan()
    n, err := strconv.ParseInt(scanner.Text(), 10, 64)
    check(err)
    var list Lines
    for scanner.Scan() {
        list = list.LongestLines(scanner.Text())
        list = list.ReduceLengthIfNeeded(n)
    }
    for _,s := range list {
        fmt.Println(s)
    }
}
