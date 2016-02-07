package main

import (
   "fmt"
   "bufio"
   "os"
   "strconv"
)

func check(err error) {
   if err != nil {
      panic(err)
   }
}

func NthFib(n int64) int64 {
   if n == 1 {
      return 1
   }
   if n == 0 {
      return 0
   }
   return NthFib(n -1) + NthFib(n - 2)
}

func main() {
   f, err := os.Open(os.Args[1])
   check(err)
   s := bufio.NewScanner(f)
   for s.Scan() {
      n, err := strconv.ParseInt(s.Text(), 10, 64)
      check(err)
      fmt.Println(NthFib(n))
   }
}
