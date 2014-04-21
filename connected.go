package main

import (
    "encoding/json"
    "log"
    "net/http"
)

func connect(rw http.ResponseWriter, req *http.Request) {
  decoder := json.NewDecoder(req.Body)
  log.Println(serving)
    //
    // req.ParseForm()
    // log.Println(req.Form)
    // //LOG: map[{"test": "that"}:[]]
    // var t test_struct
    // for key, _ := range req.Form {
    //     log.Println(key)
    //     //LOG: {"test": "that"}
    //     err := json.Unmarshal([]byte(key), &t)
    //     if err != nil {
    //         log.Println(err.Error())
    //     }
    // }
    // log.Println(t.Test)
    // //LOG: that
}

func main() {
    http.HandleFunc("/connect",connect)
    log.Fatal(http.ListenAndServe(":8080",nil))
}