# UDP packet comparator


## Project creation

Change to an empty working directory of your choice.
Then manually create the directory structure like this:
 

## Defaults
```
 TIMEOUT=2
 PORT=20000
```

## Exit codes
```
 1: Unknown application error, see output
 2: Timeout
 3: Received packet does not match expected data
```

## Run 
 
You can now run the service with docker (note the -i parameter):
```
 docker build -t udp-packet-comparator:latest .
 echo test | docker run -e PORT=20000 -e TIMEOUT=20 -i -p 20000:20000/udp udp-packet-comparator:latest   

```

Run a Test-UDP call
```
 echo test | netcat -u localhost 20000
``` 
