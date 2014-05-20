package main

import (
  "net/http"
  "github.com/gorilla/mux"
  "encoding/json"
  "io/ioutil"
  "strconv"
  "fmt"
)

type Object struct {
  P string
  Q string
}

func ObjectGetHandler(rw http.ResponseWriter,req *http.Request) {
  logger.Printf("[ObjectGetHandler] called")
  rw.Header().Set("Access-Control-Allow-Origin", "*")
  vars := mux.Vars(req)
  logger.Printf("[ObjectGetHandler] vars: %s",vars)

  var hsh map[string][]Object
  var objects []Object
  hsh = make(map[string][]Object)
  for p,q := range ids {
    objects = append(objects,Object{strconv.Itoa(p),strconv.Itoa(q)})
  }
  hsh["objects"] = objects

  response,err := json.Marshal(&hsh)

  if err != nil {
    logger.Panic("Cannot marshal!")
  }

  logger.Printf("[ObjectGetHandler] %s",response)

  fmt.Fprintf(rw,"{\"object\": { \"id\":\"2\" } }")

  logger.Printf("[ObjectGetHandler] finished")
}

func IsConnectedGetHandler(rw http.ResponseWriter,req *http.Request) {
  logger.Printf("[IsConnectedGetHandler] called")
  vars := mux.Vars(req)
  p,err := strconv.ParseInt(vars["p"],10,0)
  if err != nil {
    logger.Panic("[IsConnectedGetHandler] error converting p")
  }
  q,err := strconv.ParseInt(vars["q"],10,0)
  if err != nil {
    logger.Panic("[IsConnectedGetHandler] error converting q")
  }

  pair := Pair{int(p),int(q)}
  is_connected := connected(pair,ids)
  logger.Printf("[IsConnectedGetHandler] connected result: %s",is_connected)
  logger.Printf("[IsConnectedGetHandler] vars: %s",vars)
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
