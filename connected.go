package main

import (
    // "encoding/json"
    "fmt"
    "net/http"
    "io/ioutil"
)

func connect(rw http.ResponseWriter,req *http.Request) {
  // decoder := json.NewDecoder(req.Body)
  body, _ := ioutil.ReadAll(req.Body)
  fmt.Fprintf(rw,"parsed: "+string(body))
}

func main() {
  http.HandleFunc("/connect",connect)
  http.ListenAndServe(":9091",nil)
}