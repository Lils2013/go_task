package main

import (
  "bufio"
  "fmt"
  "os"
)

func main() {
  if len(os.Args) != 2 {
    fmt.Println("Wrong amount of arguments!")
    return
  }
  file, err := os.Open(os.Args[1])
  if err != nil {
    panic(err)
  }
  defer file.Close()
  lineCount := 0
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lineCount++
  }
  fmt.Println(lineCount, os.Args[1])
}
