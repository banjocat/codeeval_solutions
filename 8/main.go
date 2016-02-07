
package main

import (
   "fmt"
   "os"
   "bufio"
   "strings"
)

func check(err error) {
   if err != nil {
      panic(err)
   }
}


func ReverseLine(line string) string {
   scanner := bufio.NewScanner(strings.NewReader(line))
   scanner.Split(bufio.ScanWords)
   scanner.Scan()
   reverse := scanner.Text()
   for scanner.Scan() {
      reverse = scanner.Text() + " " + reverse
   }
   return reverse
}

func main() {
   file, err := os.Open(os.Args[1])
   check(err)
   scanner := bufio.NewScanner(file)
   for scanner.Scan() {
      fmt.Println(ReverseLine(scanner.Text()))
   }
}
