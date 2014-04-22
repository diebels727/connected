package main

import (
  "fmt"
  "time"
)

type Pair struct {
  p int
  q int
}

func find(p int,id []int) int {
  return id[p]
}

func connected(pair Pair,id []int) bool {
  return id[pair.p] == id[pair.q]
}

func union(p chan Pair,id []int) []int {
  pair := <- p
  if (connected(pair,id)) { return id };
  pid := id[pair.p];
  for i:=0;i<len(id);i++ {
      if (id[i] == pid) { id[i] = id[pair.q] };
  }
  return id
}

func main() {
  id := make([]int,0)
  for i:=0;i<10;i++ {
    id = append(id,i)
  }
  pchan := make(chan Pair,0)
  go union(pchan,id)
  p := Pair{1,2}
  pchan <- p

  time.Sleep(1e9)

  // id = union(3,4,id)
  // id = union(2,3,id)
  // id = union(4,5,id)
  // id = union(6,7,id)
  // id = union(0,7,id)
  // id = union(9,8,id)
  // id = union(7,1,id)
  fmt.Println(id)
}

