package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "io/ioutil"
)

type Message struct {
    P string
    Q string
}

func connect(rw http.ResponseWriter,req *http.Request) {
  // decoder := json.NewDecoder(req.Body)
  var m Message
  body, _ := ioutil.ReadAll(req.Body)
  err := json.Unmarshal(body,&m)
  if err != nil {
    fmt.Fprintf(rw,"error")
  }
  fmt.Fprintf(rw,"P: %s\nQ: %s\n",m.P,m.Q)
}

func main() {
  http.HandleFunc("/connect",connect)
  http.ListenAndServe(":9091",nil)
}