package main

import (
  "fmt"
)

func find(p int,id []int) int {
  return id[p]
}

func connected(p int,q int,id []int) bool {
  return id[p] == id[q]
}

func union(p int,q int,id []int) []int {
  if (connected(p,q,id)) { return id };
  pid := id[p];
  for i:=0;i<len(id);i++ {
      if (id[i] == pid) { id[i] = id[q] };
  }
  return id
}

func main() {
  id := make([]int,0)
  for i:=0;i<10;i++ {
    id = append(id,i)
  }
  id = union(1,2,id)
  id = union(3,4,id)
  id = union(2,3,id)
  id = union(4,5,id)
  id = union(6,7,id)
  id = union(0,7,id)
  id = union(9,8,id)
  id = union(7,1,id)
  fmt.Println(id)
}

