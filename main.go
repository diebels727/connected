package main

import (
  "net/http"
)

func main() {
  id := make([]int,0)
  for i:=0;i<10;i++ {
    id = append(id,i)
  }
  pchan := make(chan Pair,0)
  go union(pchan,id)

  http.HandleFunc("/connect",Connect)
  http.ListenAndServe(":9091",nil)
}