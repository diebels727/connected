package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "io/ioutil"
)

type Message struct {
    p string
    q string
}

func connect(rw http.ResponseWriter,req *http.Request) {
  // decoder := json.NewDecoder(req.Body)
  var m Message
  body, _ := ioutil.ReadAll(req.Body)
  err := json.Unmarshal(body,&m)
  if err != nil {
    fmt.Fprintf(rw,"error")
  }
  fmt.Fprintf(rw,"p: %s\nq: %s\n",m.p,m.q)
}

func main() {
  http.HandleFunc("/connect",connect)
  http.ListenAndServe(":9091",nil)
}