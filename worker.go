package main

import (
"os/exec"
"regexp"
"strconv"
"strings"
)

type jobInfo struct{
	Action string
	Host string
	Rc string
}

type pingReturn struct{
  Status string `json:"status"`
  Message string `json:"message,omitempty"`
  Host string `json:"host,omitempty"`
  Stats *StatsJ `json:"stats,omitempty"`
  Rtt *rttJ `json:"rtt,omitempty"`
  Logs []pingLogJ `json:"ping_logs,omitempty"`
  MtrLogs []mtrJ `json:"mtr_logs,omitempty"`
}

type rttJ struct {
  Min  string `json:"min"`
  Avg string `json:"avg"`
  Max  string `json:"max"`
  Mdev string `json:"mdev"`
}

type StatsJ struct {
  Transmitted  int `json:"transmitted"`
  Received int `json:"received"`
  Loss  string `json:"loss"`
  Time int `json:"time"`
}

type pingLogJ struct {
  Id  int `json:"id"`
  Size  int `json:"size"`
  Host string `json:"host"`
  Ttl int `json:"ttl"`
  Time string `json:"time"`
}

type mtrJ struct {
  Id  int `json:"id"`
  Host string `json:"host"`
  Loss string `json:"loss"`
  Sent int `json:"sent"`
  Last string `json:"last"`
  Avg string `json:"average"`
  Best string `json:"best"`
  Wrst string `json:"worst"`
  StDev string `json:"st_deviation"`
}

var rttR = regexp.MustCompile(`rtt min\/avg\/max\/mdev = (.*)\/(.*)\/(.*)\/(.*) ms`)
var statsR = regexp.MustCompile(`(\d*) packets transmitted, (\d*) received, (\d*)% packet loss, time (\d*)ms`)
var logR = regexp.MustCompile(`(.+) bytes from (.+): icmp_[sq]eq=(.+) ttl=(.+) time=(.+) ms`)
var mtrR = regexp.MustCompile(`\S+.\|-- (\S+)\s*(\S+)%?\s*(\d+)\s*(\S+)\s*(\S+)\s*(\S+)\s*(\S+)\s*(\S+)`)

func worker(ji jobInfo) (*pingReturn) {
  if ji.Host == ""{
    return &pingReturn{
      Status: "error",
      Message: "Host is needed",
    }
  }

  _, err := strconv.Atoi(ji.Rc)
  if err != nil{
    return &pingReturn{
      Status: "error",
      Message: "rc needs to be an integer",
    }
  }

  switch {
  case ji.Action == "ping":
    return ping(ji.Host, ji.Rc)
  case ji.Action == "mtr":
    return mtr(ji.Host, ji.Rc)
  case ji.Action == "":
    return &pingReturn{
      Status: "error",
      Message: "Action is needed",
    }
  }
  return &pingReturn{
    Status: "error",
    Message: "Action not found",
  }

}


func ping(host string, rc string) (*pingReturn) {
  allStructOut := &pingReturn {
    Host: host,
  }

//Run the ping command
  out, err := exec.Command("ping", "-c", rc, host, "-i", ".3").Output()
  if err != nil {
    return &pingReturn{
      Host: host,
      Status: "error",
      Message: err.Error(),
    }
  }

  // Get the rtt info from the output
  rttRes:= rttR.FindStringSubmatch(string(out))
  if rttRes != nil {
    allStructOut.Rtt = &rttJ{ rttRes[1], rttRes[2], rttRes[3], rttRes[4]}
  }

  // Gets the stats from the output
  statsRes:= statsR.FindStringSubmatch(string(out))
  if statsRes != nil {
    trans, err := strconv.Atoi(statsRes[1])
    recv, err := strconv.Atoi(statsRes[2])
    timeTo, err := strconv.Atoi(statsRes[4])
    if err != nil {
      return &pingReturn{
        Host: host,
        Status: "error",
        Message: err.Error(),
      }
    }
    allStructOut.Stats = &StatsJ{ trans, recv, statsRes[3], timeTo}
  }

  // Gets all the ping logs
  logRes:= logR.FindAllStringSubmatch(string(out), -1)
	for _, v := range logRes {
    iBytes, err := strconv.Atoi(v[1])
    iReq, err := strconv.Atoi(v[3])
    iTtl, err := strconv.Atoi(v[4])
    if err != nil {
      return &pingReturn{
        Host: host,
        Status: "error",
        Message: err.Error(),
      }
    }
    allStructOut.Logs = append(allStructOut.Logs, pingLogJ{iReq, iBytes, v[2], iTtl, v[5]})
  }

  allStructOut.Status = "ok"
  return allStructOut

}

//mtr runs mtr against the host
func mtr(host string, rc string) (*pingReturn) {
  allStructOut := &pingReturn {
    Host: host,
  }

  out, err := exec.Command("mtr", host, "-c", rc, "-r", "-n").Output()
  if err != nil {
    return &pingReturn{
      Host: host,
      Status: "error",
      Message: err.Error(),
    }
  }

  mtrParse:= mtrR.FindAllStringSubmatch(string(out), -1)
  for k, v := range mtrParse {

    iSent, err := strconv.Atoi(v[3])
    allStructOut.MtrLogs = append(allStructOut.MtrLogs, mtrJ{k, v[1], strings.Replace(v[2], "%", "", 1), iSent, v[4], v[5], v[6], v[7], v[8]})
    if err != nil {
      return &pingReturn{
        Host: host,
        Status: "error",
        Message: err.Error(),
      }
    }
  }

  allStructOut.Status = "ok"
  return allStructOut
}
