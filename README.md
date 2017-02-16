# gonxapi
Cisco NX-API client library written in Go

1. clone the project:
$ git clone https://github.com/danehans/gonxapi.git

2. cd into the project dir.

2. Build the bin:
$ go build -i

3. Run the client against the nx-api, passing the appropriate flags:
$ ./gonxapi -address=10.1.1.254 -username=admin -password=cisco 
