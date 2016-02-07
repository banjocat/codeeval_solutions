
package main

import (
   "fmt"
   "strings"
   "bufio"
   "os"
)

func check(err error) {
   if err != nil {
      panic(err)
   }
}

func SwapNumbers(s string) string {
   var replace = []rune(s)
   var original = []rune(s)
   replace[0] = original[len(original)-1]
   replace[len(original)-1] = original[0]
   return string(replace)
}

func HandleLine(line string) string {
   var result string
   s := bufio.NewScanner(strings.NewReader(line))
   s.Split(bufio.ScanWords)
   s.Scan()
   result += SwapNumbers(s.Text())
   for s.Scan() {
      result += " " + SwapNumbers(s.Text())
   }
   return result
}

func main() {
   f, e := os.Open(os.Args[1])
   check(e)
   s := bufio.NewScanner(f)
   for s.Scan() {
      fmt.Println(HandleLine(s.Text()))
   }
}
