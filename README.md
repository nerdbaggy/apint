# apint
apint - Use common network tools over an API and get results via json

apint returns the full output of Ping and MTR into json format with parameters.

This is something I created in my spare time to learn GO. The code is bad and I do no parameter validation. ** Do not use on open internet**

# How to Install
## Install MTR
```
apt-get install mtr
```

#### Download the proper release for your operating system
```
64 Bit
wget https://github.com/nerdbaggy/apint/releases/download/3/apint_64 -O apint

32 Bit
wget https://github.com/nerdbaggy/apint/releases/download/3/apint_32 -O apint

Arm
wget https://github.com/nerdbaggy/apint/releases/download/3/apint_arm -O apint
```
#### Start apint
```
./apint
```

# How to Use
Communication with apint is done via a web interface. This isn't RESTFUL its just what I decided to use.

Either a POST or GET request can be used to get the results. POST sends a json payload while GET uses the URL parameters

## Parameters
**action** - [ping, mtr] - Either ping the host or run MTR against the host<br>
**host** - The host to interact with<br>
**rc** How many requests to send<br>
**callback** - Used with jsonp
## POST
```
curl -X POST -d '{"action": "ping", "host": "8.8.8.8", "rc": "3"}' http://localhost:8080
```
#### Format - Ping
```
{
	"action": "ping",
	"host": "8.8.8.8",
	"rc": "3"
}
```
#### Format - MTR
```
{
	"action": "mtr",
	"host": "8.8.8.8",
	"rc": "3"
}
```

## GET
```
curl "localhost:8080?action=ping&host=8.8.8.8&rc=3"
```

# Response
## Ping ##
```
{
	"status": "ok",
	"host": "8.8.8.8",
	"stats": {
		"transmitted": 3,
		"received": 3,
		"loss": "0",
		"time": 603
	},
	"rtt": {
		"min": "33.085",
		"avg": "37.271",
		"max": "44.522",
		"mdev": "5.147"
	},
	"logs": [{
		"id": 1,
		"size": 64,
		"host": "8.8.8.8",
		"ttl": 63,
		"time": "44.5"
	}, {
		"id": 2,
		"size": 64,
		"host": "8.8.8.8",
		"ttl": 63,
		"time": "33.0"
	}, {
		"id": 3,
		"size": 64,
		"host": "8.8.8.8",
		"ttl": 63,
		"time": "34.2"
	}]
}
```

## MTR
```
{
	"status": "ok",
	"host": "8.8.8.8",
	"logs": [{
		"id": 7,
		"host": "68.86.91.25",
		"loss": "0.0",
		"sent": 3,
		"last": "27.9",
		"average": "27.0",
		"best": "26.4",
		"worst": "27.9",
		"st_deviation": "0.8"
	}, {
		"id": 8,
		"host": "68.86.87.142",
		"loss": "0.0",
		"sent": 3,
		"last": "27.2",
		"average": "25.8",
		"best": "24.8",
		"worst": "27.2",
		"st_deviation": "1.2"
	}, {
		"id": 9,
		"host": "75.149.229.86",
		"loss": "0.0",
		"sent": 3,
		"last": "29.5",
		"average": "42.3",
		"best": "29.5",
		"worst": "53.8",
		"st_deviation": "12.2"
	}, {
		"id": 10,
		"host": "209.85.242.142",
		"loss": "0.0",
		"sent": 3,
		"last": "25.3",
		"average": "27.5",
		"best": "25.1",
		"worst": "32.2",
		"st_deviation": "4.0"
	}, {
		"id": 11,
		"host": "72.14.236.152",
		"loss": "0.0",
		"sent": 3,
		"last": "41.3",
		"average": "37.3",
		"best": "26.7",
		"worst": "44.1",
		"st_deviation": "9.3"
	}, {
		"id": 12,
		"host": "72.14.235.10",
		"loss": "0.0",
		"sent": 3,
		"last": "33.3",
		"average": "34.8",
		"best": "33.3",
		"worst": "37.3",
		"st_deviation": "2.2"
	}, {
		"id": 13,
		"host": "72.14.234.53",
		"loss": "0.0",
		"sent": 3,
		"last": "41.2",
		"average": "40.2",
		"best": "35.5",
		"worst": "43.7",
		"st_deviation": "4.2"
	}, {
		"id": 14,
		"host": "???",
		"loss": "100.0",
		"sent": 3,
		"last": "0.0",
		"average": "0.0",
		"best": "0.0",
		"worst": "0.0",
		"st_deviation": "0.0"
	}, {
		"id": 15,
		"host": "8.8.8.8",
		"loss": "0.0",
		"sent": 3,
		"last": "35.0",
		"average": "34.3",
		"best": "33.3",
		"worst": "35.0",
		"st_deviation": "0.9"
	}]
}
```

# Errors
**200** - Everything worked fine<br>
**500** - Everything else that broke

**Status** - will always equal "ok" if everything worked correctly. If not it will display "error"<br>
**Message** - Displays the error that occurred to caused a 500 response
