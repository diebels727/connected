package main

import (
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
    logger.Printf("[main] set initial condition for ready state")
    ready <- true
    logger.Printf("[main] Launching main event loop")
    for {
      logger.Printf("[main] Waiting to pull pair from channel ...")
      p := <- pchan
      logger.Printf("[main] Pulled pair from channel")
      union(p,id)
      logger.Printf("[main] Executed union")
    }
    logger.Printf("Terminating main event loop ...")
  }();

  http.HandleFunc("/connect",Connect)
  http.HandleFunc("/connected",Connected)
  http.HandleFunc("/object",ObjectHandler)
  http.ListenAndServe(":9091",nil)
}