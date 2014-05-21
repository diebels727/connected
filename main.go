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

var ids []int
var ready chan bool
var pchan chan Pair
var log_file *os.File
var logger *log.Logger

func main() {
  log_file,_ = os.Create("connected.log")
  logger = log.New(log_file,"",LOG_FORMAT)
  logger.Printf("Starting up ...")

  ids = make([]int,0)
  for i:=0;i<10;i++ {
    ids = append(ids,i)
  }
  ready = make(chan bool,0)
  pchan = make(chan Pair,0)

  go func(){
    logger.Printf("[main] Launching main event loop")
    for {
      p := <- pchan
      union(p,ids)
    }
    logger.Printf("Terminating main event loop ...")
  }();

  m := mux.NewRouter()
  m.HandleFunc("/records",GetRecordsHandler).Methods("GET")
  // m.HandleFunc("/records",PostRecordsHandler).Methods("POST")
  m.HandleFunc("/records",OptionsRecordsHandler).Methods("OPTIONS")
  m.HandleFunc("/records/{id}",ShowRecordsHandler).Methods("GET")

  http.ListenAndServe(":9091",m)
}