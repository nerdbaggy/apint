package main

import (
"strconv"
"testing"
)

func TestFibFunc(t *testing.T) {

	var jobTest jobInfo
	jobTest.Action = "ping"
	jobTest.Host = "127.0.0.1"
	jobTest.Rc = "3"

	rcInt, err := strconv.Atoi(jobTest.Rc)

	if err != nil{
		t.Errorf("Error parsing rc: %s", err)
	}

    fn := worker(jobTest)

    if fn.Status != "ok" {
    	t.Errorf("ok status not returned, got %s", fn.Status)
    }

    if fn.Message != "" {
    	t.Errorf("Message was returned when it should not of been, got %s", fn.Message)
    }

    if fn.Host != jobTest.Host {
    	t.Errorf("Host was not returned with the same as sent, got %s", fn.Host)
    }

    if fn.Stats.Transmitted != rcInt {
    	t.Errorf("Incorrect number of packets sent, got %d", fn.Stats.Transmitted)
    }

    if fn.Stats.Received != rcInt {
    	t.Errorf("Incorrect number of packets recieved, got %d", fn.Stats.Received)
    }

    if fn.Stats.Loss != "0" {
    	t.Errorf("Incorrect packet loss, got %s", fn.Stats.Loss)
    }

    if fn.Stats.Time == rcInt {
    	t.Errorf("Total ping time was 0")
    }

    if fn.Rtt.Min == ""{
    	t.Errorf("min returned empty")
    }

    if fn.Rtt.Avg == ""{
    	t.Errorf("avg returned empty")
    }

    if fn.Rtt.Max == ""{
    	t.Errorf("max returned empty")
    }

    if fn.Rtt.Mdev == ""{
    	t.Errorf("mdev returned empty")
    }

    if len(fn.Logs) != rcInt{
    	t.Errorf("Expecting 3 logs back, got %d", len(fn.Logs))
    }

    for k, v := range fn.Logs {
    	if k+1 != v.Id{
    		t.Errorf("Incorrect log id, got %d", v.Id)
    	}

    	if v.Size != 64{
    		t.Errorf("Incorrect packet size, got %d", v.Size)
    	}

    	if v.Host != jobTest.Host {
    		t.Errorf("Host was not returned with the same as sent, got %s", v.Host)
   		}

   		if v.Ttl == 0 {
    		t.Errorf("TTL not recieved")
   		}

   		if v.Time == ""  {
    		t.Errorf("Time not recieved")
   		}

    }
}