package main

import (
  "net/http"
)

var pchan chan Pair
var id []int

func main() {
  id = make([]int,0)
  for i:=0;i<10;i++ {
    id = append(id,i)
  }
  pchan = make(chan Pair,0)

  go func(){
    for {
      p := <- pchan
      union(p,id)
    }
  }();

  http.HandleFunc("/connect",Connect)
  http.ListenAndServe(":9091",nil)
}