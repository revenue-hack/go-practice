package main

import (
  "fmt"
  "ch02/ex01/tempconv"
  "os"
  "strconv"
)

func main() {
  for _, arg := range os.Args[1:] {
    t, err := strconv.ParseFloat(arg, 64)
    if err != nil {
      fmt.Fprintf(os.Stderr, "cf: %v\n", err)
      os.Exit(2)
    }
    f := tempconv.Fahrenheit(t)
    c := tempconv.Celsius(t)
    fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
  }
}

