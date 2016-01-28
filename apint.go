package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"flag"
)


func handler(w http.ResponseWriter, r *http.Request) {
  var jobPost jobInfo
  if r.Method == "POST" {
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&jobPost)
    if err != nil {
      fmt.Println(err)
    }
    } else if r.Method == "GET" {
      jobPost.Action = r.URL.Query().Get("action")
      jobPost.Host = r.URL.Query().Get("host")
      jobPost.Rc = r.URL.Query().Get("rc")
    }

    resp := worker(jobPost)
    jsonResponse, _ := json.Marshal(resp)
    w.Header().Set("Access-Control-Allow-Origin", "*")


    if resp.Status != "error"{
      if r.URL.Query().Get("callback") != ""{
        w.Header().Set("Content-Type", "application/javascript")
        fmt.Fprintf(w, "%s(%s);", r.URL.Query().Get("callback"), jsonResponse)
      } else{
        w.Header().Set("Content-Type", "application/json")
        fmt.Fprintf(w, "%s", jsonResponse)
      }
    } else{
      w.WriteHeader(http.StatusInternalServerError)
      if r.URL.Query().Get("callback") != ""{
        w.Header().Set("Content-Type", "application/javascript")
        fmt.Fprintf(w, "%s(%s);", r.URL.Query().Get("callback"), jsonResponse)
      } else{
        w.Header().Set("Content-Type", "application/json")

        fmt.Fprintf(w, "%s", jsonResponse)
      }
    }

  }

  func main() {
    port := flag.Int("port", 8080, "Port that the web server will listen on. Default is 8080")
    ip := flag.String("ip", "", "IP to listen on. Default is all ips")
    flag.Parse()

    http.HandleFunc("/", handler)
    http.ListenAndServe(fmt.Sprintf("%s:%d", *ip, *port), nil)
  }
