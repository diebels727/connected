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
    for {
      p := <- pchan
      union(p,id)
    }
  }();

  http.HandleFunc("/connect",Connect)
  http.HandleFunc("/connected",Connected)
  http.ListenAndServe(":9091",nil)
}