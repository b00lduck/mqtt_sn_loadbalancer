# mqtt_sn_loadbalancer

## Golang general

 * Tutorial (go official): https://tour.golang.org/welcome/1
 * Reference (go official): https://golang.org/doc/effective_go.html

## Frameworks

 * Configuration: t.b.d.
 * Logging: https://github.com/Sirupsen/logrus
 * MongoDB: https://github.com/go-mgo/mgo/tree/v2
 * HAProxyV2: https://godoc.org/github.com/pires/go-proxyproto

Useful Go built-in features:
 * HTTP server (https://golang.org/pkg/net/http/)
 * UDP server (https://golang.org/pkg/net/)
 * Test framework (https://golang.org/pkg/testing/)
 * Test runner (go test)
 * Static code analyis (govet)
  

## Project creation

Change to an empty working directory of your choice.
Then manually create the directory structure like this:
 
```
 mkdir gopath
 cd gopath
 export GOPATH=$(pwd)
 mkdir -p src/github.com/b00lduck
 cd src/github.com/b00lduck
 git clone git@github.com:b00lduck/mqtt_sn_loadbalancer.git
 cd mqtt_sn_loadbalancer
 go get ./...
```
   
The project is set up and your GOPATH is now set. Check this with:

 echo $GOPATH
 
## Run 
 
You can now run the service directly with go:
```
 go run . 
```

or with docker:
```
 docker build -t mqtt_sn_loadbalancer:latest . && docker run mqtt_sn_loadbalancer:latest 
```

Run a Test-UDP call (misusing nslookup as UDP-generator ;D)with:
```
 nslookup -port=2000 blah localhost
``` 

## Running Tests 
 
To run the tests call:
```
 go test ./...
``` 
