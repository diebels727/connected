package main

import (
  "net/http"
  "github.com/gorilla/mux"
  "encoding/json"
  "io/ioutil"
  "strconv"
  "fmt"
  "strings"
)

type Record struct {
  Id string
  Value string
}

type ParseRecord struct {
  Record Record
}

func GetRecordsHandler(rw http.ResponseWriter,req *http.Request) {
  logger.Printf("[GetRecords] called")
  rw.Header().Set("Access-Control-Allow-Origin", "*")

  var records_with_root map[string][]Record
  records_with_root = make(map[string][]Record)

  var records []Record

  for id,value := range ids {
    records = append(records,Record{strconv.Itoa(id),strconv.Itoa(value)})
  }
  records_with_root["records"] = records
  response,err := json.Marshal(&records_with_root)
  if err != nil {
    logger.Panic("[GetRecords] Cannot marshal records.")
  }

  response_string := strings.ToLower(string(response))

  fmt.Fprintf(rw,response_string) //output to consumer

  logger.Printf("[GetRecords] finished")
}

func ShowRecordsHandler(rw http.ResponseWriter,req *http.Request) {
  logger.Printf("[ShowRecords] called")
  rw.Header().Set("Access-Control-Allow-Origin", "*")
  var record_with_root map[string]Record
  record_with_root = make(map[string]Record)
  vars := mux.Vars(req)
  id_string := vars["id"]
  id,err := strconv.Atoi(id_string)
  if err != nil {
    logger.Panic("[ShowRecords] Cannot convert id: %s",err)
  }
  value := ids[id]
  record := Record{strconv.Itoa(id),strconv.Itoa(value)}
  record_with_root["record"] = record
  response,err := json.Marshal(&record_with_root)
  if err != nil {
    logger.Panic("[ShowRecords] Cannot marshal record.")
  }
  response_string := strings.ToLower(string(response))
  fmt.Fprintf(rw,response_string) //output to consumer
  logger.Printf("[ShowRecords] finished")
}

func PostRecordsHandler(rw http.ResponseWriter,req *http.Request) {
  logger.Printf("[PostRecordsHandler] called")
  rw.Header().Set("Access-Control-Allow-Headers","Origin, X-Requested-With, Content-Type, Accept")
  rw.Header().Set("Access-Control-Allow-Origin", "*")
  //TODO: localhost should not be statically defined; but this works for localhost testing
  rw.Header().Set("Origin","http://localhost:9091")
  var parse_record ParseRecord
  body, _ := ioutil.ReadAll(req.Body)
  err := json.Unmarshal(body,&parse_record)
  if err != nil {
    logger.Panic("[PostRecordsHandler] error in JSON: %s",err)
  }
  id,err := strconv.ParseInt(parse_record.Record.Id,10,0)
  if err != nil {
    logger.Panic("[PostRecordsHandler] error converting id: %s",err)
  }
  value,err := strconv.ParseInt(parse_record.Record.Value,10,0)
  if err != nil {
    logger.Panic("[PostRecordsHandler] error converting value: %s",err)
  }
  pair := Pair{int(id),int(value)}
  logger.Printf("[PostRecordsHandler] parsed id: %s and value: %s",id,value)
  logger.Printf("[PostRecordsHandler] sending pair on channel ...")
  pchan <- pair
  logger.Printf("[PostRecordsHandler] sent pair on channel")

  fmt.Fprintf(rw,"{}") //empty content response; ember seemed to complain

  logger.Printf("[TempPostHandler] finished")
}

func OptionsRecordsHandler(rw http.ResponseWriter,req *http.Request) {
  logger.Printf("[OptionsRecordsHandler] called")
  rw.Header().Set("Access-Control-Allow-Origin", "*")
  rw.Header().Set("Access-Control-Allow-Headers","Origin, X-Requested-With, Content-Type, Accept")
  logger.Printf("[OptionsRecordsHandler] %s",rw.Header())
}
