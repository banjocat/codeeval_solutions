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

func main() {
   f, e := os.Open(os.Args[1])
   check(e)
   s := bufio.NewScanner(f)
   var sum int64
   for s.Scan() {
      n, e := strconv.ParseInt(s.Text(), 10, 64)
      check(e)
      sum += n
   }
   fmt.Println(sum)
}
