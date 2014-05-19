package main

import (
  "net/http"
  "github.com/gorilla/mux"
  // "encoding/json"
  // "io/ioutil"
  // "fmt"
  // "strconv"
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
}
