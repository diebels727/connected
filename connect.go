package main

import (
  "net/http"
  "encoding/json"
  "io/ioutil"
  "fmt"
)

type Connection struct {
    P string
    Q string
}

func Connect(rw http.ResponseWriter,req *http.Request) {
  var o Connection
  body, _ := ioutil.ReadAll(req.Body)
  err := json.Unmarshal(body,&o)
  if err != nil {
    fmt.Fprintf(rw,"error")
  }
  fmt.Fprintf(rw,"P: %s\nQ: %s\n",o.P,o.Q)
}
