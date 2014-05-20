package main

// import (
//   "fmt"
//   "time"
// )

type Pair struct {
  P int
  Q int
}

func find(p int,ids []int) int {
  return ids[p]
}

func connected(pair Pair,ids []int) bool {
  return ids[pair.P] == ids[pair.Q]
}

func union(pair Pair,ids []int) []int {
  logger.Printf("[union] called")
  if (connected(pair,ids)) { return ids };
  pid := ids[pair.P];
  for i:=0;i<len(ids);i++ {
      if (ids[i] == pid) { ids[i] = ids[pair.Q] };
  }
  logger.Printf("[union] finished")
  return ids
}
