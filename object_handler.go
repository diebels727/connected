package main

import (
  "net/http"
  "github.com/gorilla/mux"
  "encoding/json"
  "io/ioutil"
  "strconv"
  "fmt"
  "strings"
)

type Object struct {
  P string
  Q string
}

type ParseObject struct {
  Object Object
}

type ResponseObject struct {
  Id string
  P string
}

func TempObjectsGetHandler(rw http.ResponseWriter,req *http.Request) {
  logger.Printf("[TempObjectsGetHandler] called")

  logger.Printf("[TempObjectsGetHandler] finished")
}

func ObjectsGetHandler(rw http.ResponseWriter,req *http.Request) {
  logger.Printf("[ObjectGetHandler] called")
  rw.Header().Set("Access-Control-Allow-Origin", "*")

  //Ugly, but working
  var hsh map[string][]ResponseObject
  var objects []ResponseObject
  hsh = make(map[string][]ResponseObject)
  for id,p := range ids {
    objects = append(objects,ResponseObject{strconv.Itoa(id),strconv.Itoa(p)})
  }
  hsh["objects"] = objects

  response,err := json.Marshal(&hsh)
  if err != nil {
    logger.Panic("Cannot marshal!")
  }
  rsp := strings.ToLower(string(response))
  //end ugly

  fmt.Fprintf(rw,rsp)
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

func TempPostHandler(rw http.ResponseWriter,req *http.Request) {
  logger.Printf("[TempPostHandler] called")
  rw.Header().Set("Access-Control-Allow-Headers","Origin, X-Requested-With, Content-Type, Accept")
  rw.Header().Set("Access-Control-Allow-Origin", "*")
  rw.Header().Set("Origin","http://localhost:9091")
  var parse_object ParseObject
  body, _ := ioutil.ReadAll(req.Body)
  err := json.Unmarshal(body,&parse_object)
  if err != nil {
    logger.Panic("[TempPostHandler] error in JSON")
  }
  p,err := strconv.ParseInt(parse_object.Object.P,10,0)
  if err != nil {
    logger.Panic("[TempPostHandler] error converting p: %s",err)
  }
  q,err := strconv.ParseInt(parse_object.Object.Q,10,0)
  if err != nil {
    logger.Panic("[TempPostHandler] error converting q: %s",err)
  }
  pair := Pair{int(p),int(q)}
  logger.Printf("[TempPostHandler] parsed p: %s and q: %s",p,q)
  logger.Printf("[TempPostHandler] sending pair on channel ...")
  pchan <- pair
  logger.Printf("[TempPostHandler] sent pair on channel")

  fmt.Fprintf(rw,"{}")
  logger.Printf("[TempPostHandler] finished")
}

func OptionsPostHandler(rw http.ResponseWriter,req *http.Request) {
  logger.Printf("[OptionsPostHandler] called")
  rw.Header().Set("Access-Control-Allow-Origin", "*")
  rw.Header().Set("Access-Control-Allow-Headers","Origin, X-Requested-With, Content-Type, Accept")
  logger.Printf("[OptionsPostHandler] %s",rw.Header())
}


func ObjectPostHandler(rw http.ResponseWriter,req *http.Request) {
  logger.Printf("[ObjectPostHandler] called")
  rw.Header().Set("Access-Control-Allow-Origin", "*")
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
