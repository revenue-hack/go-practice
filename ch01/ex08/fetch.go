package main

import(
  "fmt"
  "io/ioutil"
  "net/http"
  "os"
  "strings"
)

func main() {
  for _, url := range os.Args[1:] {
    resp, err := http.Get(prefixHttp(url))
    if err != nil {
      fmt.Fprintf(os.Stderr, "fetch %v\n", err)
      os.Exit(1)
    }
    b, err := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    if err != nil {
      fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
      os.Exit(1)
    }
    fmt.Printf("%s", b)
  }
}

func prefixHttp(url string) string {
  if strings.HasPrefix(url, "http://") {
    return url
  }
  return "http://" + url
}

