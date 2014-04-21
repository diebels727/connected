package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "io/ioutil"
)

type Object struct {
    P string
    Q string
}

func connect(rw http.ResponseWriter,req *http.Request) {
  // decoder := json.NewDecoder(req.Body)
  var o Object
  body, _ := ioutil.ReadAll(req.Body)
  err := json.Unmarshal(body,&o)
  if err != nil {
    fmt.Fprintf(rw,"error")
  }
  Connector(o,rw)
}

func main() {
  http.HandleFunc("/connect",connect)
  http.ListenAndServe(":9091",nil)
}