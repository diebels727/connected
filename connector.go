package main

import (
  "net/http"
  "fmt"
)

func Connector(object Object,rw http.ResponseWriter) {
  fmt.Fprintf(rw,"P: %s\nQ: %s\n",object.P,object.Q)
}