# mqtt_sn_loadbalancer

## Project creation

Change to an empty working directory of your choice.
Then manually create the directory structure like this:
 
 mkdir gopath
 cd gopath
 export $GOPATH=$(pwd)
 mkdir -p src/github.com/b00lduck
 cd src/github.com/b00lduck
 cd gopath/src/github.com/b00lduck
 git clone git@github.com:b00lduck/mqtt_sn_loadbalancer.git
 go get ./...
   
The project is set up and your GOPATH is now set. Check this with:

 echo $GOPATH
 
## Run 
 
You can now run the service with:

 go run .
 
## Test 
 
To run the tests call:

 go test ./...
 
