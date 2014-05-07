package main

import (
  "net/http"
)

var pchan chan Pair
var id []int
var ready chan bool

func main() {
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