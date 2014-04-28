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
  if (connected(pair,id)) { return id };
  pid := id[pair.P];
  for i:=0;i<len(id);i++ {
      if (id[i] == pid) { id[i] = id[pair.Q] };
  }
  return id
}

// func main() {
//   id := make([]int,0)
//   for i:=0;i<10;i++ {
//     id = append(id,i)
//   }
//
//   pchan := make(chan Pair,0)
//
//   go func(){
//     for {
//       p := <- pchan
//       fmt.Println("Chan: ",p)
//       union(p,id)
//     }
//   }();
//
//   p := Pair{1,2}
//   pchan <- p
//
//   p = Pair{1,7}
//   pchan <- p
//
//   time.Sleep(1e9)
//
//   fmt.Println(id)
// }
