package main

// import (
//   "fmt"
//   "time"
// )

type Pair struct {
  P int
  Q int
}

func find(p int,id []int) int {
  return id[p]
}

func connected(pair Pair,id []int) bool {
  return id[pair.P] == id[pair.Q]
}

func union(pair Pair,id []int) []int {
  logger.Printf("[union] called")
  if (connected(pair,id)) { return id };
  pid := id[pair.P];
  for i:=0;i<len(id);i++ {
      if (id[i] == pid) { id[i] = id[pair.Q] };
  }
  logger.Printf("[union] finished")
  return id
}
