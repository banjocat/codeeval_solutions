
package main

import (
   "fmt"
   "regexp"
   "strconv"
   "os"
   "bufio"
)

func check(err error) {
   if err != nil {
      panic(err)
   }
}

func SmallestMultiple (x, n int64) int64  {
   multiple := n
   for ; multiple < x; multiple += n {
   }
   return multiple
}

func ProcessLine(line string) int64 {
   reg := regexp.MustCompile(`(\d+),(\d+)`)
   match := reg.FindStringSubmatch(line)
   x, err := strconv.ParseInt(match[1], 10, 64)
   check(err)
   n, err := strconv.ParseInt(match[2], 10, 64)
   check(err)
   return SmallestMultiple(x, n)
}

func main() {
   file, err := os.Open(os.Args[1])
   check(err)
   scanner := bufio.NewScanner(file)
   for scanner.Scan() {
      fmt.Println(ProcessLine(scanner.Text()))
   }
}
