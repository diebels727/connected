package main

import (
  "net/http"
  "github.com/gorilla/mux"
  "encoding/json"
  "io/ioutil"
  "strconv"
)

type Object struct {
  P string
  Q string
}

func ObjectGetHandler(rw http.ResponseWriter,req *http.Request) {
  logger.Printf("[ObjectGetHandler] called")
  vars := mux.Vars(req)
  logger.Printf("[ObjectGetHandler] vars: %s",vars)
}

func ObjectPostHandler(rw http.ResponseWriter,req *http.Request) {
  logger.Printf("[ObjectPostHandler] called")
  var object Object
  body, _ := ioutil.ReadAll(req.Body)
  vars := mux.Vars(req)
  //Parse JSON
  err := json.Unmarshal(body,&object)
  if err != nil {
    logger.Panic("[ObjectPostHandler] error in JSON")
  }
  p,err := strconv.ParseInt(vars["p"],10,0)
  if err != nil {
    logger.Panic("[ObjectPostHandler] error converting p")
  }
  q,err := strconv.ParseInt(object.Q,10,0)
  if err != nil {
    logger.Panic("[ObjectPostHandler] error converting q")
  }

  pair := Pair{int(p),int(q)}

  logger.Printf("[ObjectPostHandler] sending pair on channel ...")
  pchan <- pair
  logger.Printf("[ObjectPostHandler] sent pair on channel")


  logger.Printf("[ObjectPostHandler] finished")
}
