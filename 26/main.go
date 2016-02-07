
package main

import (
   "fmt"
   "os"
)


func check(err error) {
   if err != nil {
      panic(err)
   }
}

func main() {
   f, e := os.Open(os.Args[1])
   check(e)
   s, e := f.Stat()
   check(e)
   fmt.Println(s.Size())
}
