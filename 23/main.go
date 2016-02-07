package main

import (
   "fmt"
)


func main() {
   for x := 1; x <= 12; x++ {
      for y := 1; y <= 12; y++ {
         fmt.Printf("%4d", y * x)
      }
      fmt.Print("\n")
   }
}
