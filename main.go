package main

import (
  "github.com/gorilla/mux"
  "net/http"
  "os"
  "log"
)

const (
  LOG_FORMAT = 3
)

var id []int
var ready chan bool
var pchan chan Pair
var log_file *os.File
var logger *log.Logger

func main() {
  log_file,_ = os.Create("connected.log")
  logger = log.New(log_file,"",LOG_FORMAT)
  logger.Printf("Starting up ...")

  id = make([]int,0)
  for i:=0;i<10;i++ {
    id = append(id,i)
  }
  ready = make(chan bool,0)
  pchan = make(chan Pair,0)

  go func(){
    logger.Printf("[main] Launching main event loop")
    for {
      p := <- pchan
      union(p,id)
    }
    logger.Printf("Terminating main event loop ...")
  }();

  m := mux.NewRouter()
  m.HandleFunc("/object/{p}",ObjectGetHandler).Methods("GET")
  m.HandleFunc("/object/{p}",ObjectPostHandler).Methods("POST")
  http.ListenAndServe(":9091",m)
}