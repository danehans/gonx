# gonxapi
A Cisco NX-API client library written in Go

## Switch Preperation
Make sure your nexus switch is running at least version 7.2.

Enable the nxapi feature:
`(config)# feature nxapi`

You can optionally change the listening port:
`(config)# nxapi http port 8080`

You can also enable the Cisco NX-API for HTTPS
`(config)# nxapi https port 8181`

## To use the gonx client lib:
`go get github.com/danehans/gonx/client/insapi`

## To build the command line utility:
Clone the project:
`$ git clone https://github.com/danehans/gonxapi.git`

Change into into the cmd/gonxctl project dir:
`$ cd gonx/cmd/nxctl`
and build the bin:
`$ go build -i`
