package main

import "fmt"
import "strconv"
import "os"
import "regexp"
import "bufio"

func check(e error) {
    if e != nil {
    panic(e)
    }
}


func fizzbuzz(number, fizz, buzz int64) string {

    switch {
    case number % fizz == 0 && number % buzz == 0:
        return "FB"
    case number % fizz == 0:
        return "F"
    case number % buzz == 0:
        return "B"
    }
    return strconv.FormatInt(number, 10)
}

func doOneLine(fizz, buzz, count int64) string {
    var output string
    for x := int64(1); x <= count; x++ {
        output += fizzbuzz(x, fizz, buzz)
        if x != count {
            output += " "
        }
    }
    return output
}

func parseLine (line string) {
    reg := regexp.MustCompile(`^(\d+) (\d+) (\d+)`)
    match := reg.FindStringSubmatch(line)
    fizz,err := strconv.ParseInt(match[1], 10, 64)
    check(err)
    buzz,err := strconv.ParseInt(match[2], 10, 64)
    check(err)
    numb,err := strconv.ParseInt(match[3], 10, 64)
    check(err)
    fmt.Println(doOneLine(fizz, buzz, numb))
}

func doFile(filename string) {
    f, err := os.Open(filename)
    check(err)
    defer f.Close()

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        parseLine(scanner.Text())
    }
}


func main() {
    filename := os.Args[1]
    doFile(filename)
}
