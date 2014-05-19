package main

import (
  "net/http"
  "encoding/json"
  "io/ioutil"
  "fmt"
  "strconv"
)

type Object struct {
  P string
  Q string
}

func Connect(rw http.ResponseWriter,req *http.Request) {
  var o Object
  body, _ := ioutil.ReadAll(req.Body)
  err := json.Unmarshal(body,&o)
  if err != nil {
    fmt.Fprintf(rw,"error: ")
  }

  //convert Object to Pair
  p,_ := strconv.ParseInt(o.P,10,0)
  q,_ := strconv.ParseInt(o.Q,10,0)

  pair := Pair{int(p),int(q)}

  pchan <- pair

  fmt.Fprintf(rw,"P: %d\nQ: %d\n",p,q)
}
