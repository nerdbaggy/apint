package main

import (
"strconv"
"testing"
)

func TestWorkerFunc(t *testing.T) {
	var jobTest jobInfo
	jobTest.Action = "mtr"
	jobTest.Rc = "2"

	fn := worker(jobTest)

	if fn.Status != "error"{
		t.Errorf("Expected error, got %s", fn.Status)
	}

	if fn.Message != "Host is needed"{
		t.Errorf("Expected no host message, got %s", fn.Message)
	}

	jobTest.Host = "8.8.8.8"
	jobTest.Rc = "fsdfs"
	fn = worker(jobTest)

	if fn.Status != "error"{
		t.Errorf("Expected error, got %s", fn.Status)
	}

	if fn.Message != "rc needs to be an integer"{
		t.Errorf("Expected bad rc message, got %s", fn.Message)
	}

	jobTest.Rc = "3"
	jobTest.Action = "fdsfasdfsd"
	fn = worker(jobTest)

	if fn.Status != "error"{
		t.Errorf("Expected error, got %s", fn.Status)
	}

	if fn.Message != "Action not found"{
		t.Errorf("Expected action not found message, got %s", fn.Message)
	}

	jobTest.Action = ""
	fn = worker(jobTest)

	if fn.Status != "error"{
		t.Errorf("Expected error, got %s", fn.Status)
	}

	if fn.Message != "Action is needed"{
		t.Errorf("Expected action needed message, got %s", fn.Message)
	}

}

func TestWorkerMTRFunc(t *testing.T) {
	var jobTest jobInfo
	jobTest.Action = "mtr"
	jobTest.Host = "8.8.8.8"
	jobTest.Rc = "2"

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

	if len(fn.MtrLogs) == 0{
		t.Errorf("No logs recieved")
	}

	for k, v := range fn.MtrLogs {
		if k != v.Id {
			t.Errorf("ID does not match for number")
		}

		if v.Host == "" {
			t.Errorf("No host recieved")
		}

		if v.Loss == "" {
			t.Errorf("No lost recieved")
		}

		if v.Loss == "" {
			t.Errorf("No lost recieved")
		}

		if v.Sent != rcInt {
			t.Errorf("Number of packets sent doesn't match requested, got %d", v.Sent)
		}

		if v.Last == "" {
			t.Errorf("No last recieved")
		}

		if v.Avg == "" {
			t.Errorf("No average recieved")
		}

		if v.Best == "" {
			t.Errorf("No best recieved")
		}

		if v.Wrst == "" {
			t.Errorf("No worst recieved")
		}

		if v.StDev == "" {
			t.Errorf("No st_deviation recieved")
		}

	}

}

func TestWorkerPingFunc(t *testing.T) {

	var jobTest jobInfo
	jobTest.Action = "ping"
	jobTest.Host = "8.8.8.8"
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