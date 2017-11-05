package main
import (
  "fmt"
  "os"
)

func main () {
  for i, arg := range os.Args {
    fmt.Printf("i: %d value: %s\n", i, arg)
  }
}

