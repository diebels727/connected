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
  logger.Printf("[Connect] called ...")

  var o Object
  body, _ := ioutil.ReadAll(req.Body)
  err := json.Unmarshal(body,&o)
  if err != nil {
    fmt.Fprintf(rw,"error: ")
  }

  logger.Printf("[Connect] beginning to parse integers")

  //convert Object to Pair
  p,_ := strconv.ParseInt(o.P,10,0)
  q,_ := strconv.ParseInt(o.Q,10,0)

  logger.Printf("[Connect] parsed integers")


  pair := Pair{int(p),int(q)}

  logger.Printf("[Connect] pushing pair on to channel")
  pchan <- pair
  logger.Printf("[Connect] pushed pair on to channel")
  fmt.Fprintf(rw,"P: %d\nQ: %d\n",p,q)
}
