package client

import (
  "net/http"
  "log"
  "io/ioutil"
  "fmt"
)

func PrintResponseBody(url string) {
    res, err := http.Get(url)
    if err != nil {
      log.Fatal(err)
    }
    body, err := ioutil.ReadAll(res.Body)
    res.Body.Close()
    if err != nil {
      log.Fatal(err)
    }
    fmt.Printf("%s", body)
}