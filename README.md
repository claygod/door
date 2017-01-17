Nimble multiplexer for Fast HTTP server, written in Golang.
Door is made on the basis of the router Bxog (https://github.com/claygod/Bxog) for use with `Fast HTTP server`.

[![API documentation](https://godoc.org/github.com/claygod/door?status.svg)](https://godoc.org/github.com/claygod/door)
[![Go Report Card](https://goreportcard.com/badge/github.com/claygod/door)](https://goreportcard.com/report/github.com/claygod/door)

# Usage

An example of using the multiplexer:
```go
package main

import (
	"fmt"

	"github.com/claygod/door"
	"github.com/valyala/fasthttp"
)

// Handlers
func IHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "200 OK\r\nHello, world! Requested path is %q",
		ctx.Path())
}
func THandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "200 OK\r\nParam `:par` -  %s",
		ctx.UserValue("par"))
}
func PHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "200 OK\r\nParams:\r\n`country` -  %s\r\n`capital` -  %s\r\n`valuta` -  %s",
		ctx.UserValue("name"), ctx.UserValue("city"), ctx.UserValue("money"))
}
func FSHandler(ctx *fasthttp.RequestCtx) {
	filesHandler(ctx)
}

var (
	filesHandler = fasthttp.FSHandler("./", 0)
)

// Main
func main() {
	m := door.New()
	m.Add("/", IHandler)
	m.Add("/abc/:par", THandler)
	m.Add("/country/:name/capital/:city/valuta/:money", PHandler).
		Id("country"). // For ease indicate the short ID
		Method("GET")  // GET method do not need to write here, it is used by default (this is an example)
	m.Add("/file/", FSHandler)

	m.Start(":80")
}
```

Click URLs:
- http://localhost
- http://localhost/abc/123
- http://localhost/country/USA/capital/Washington/valuta/dollar

# Settings

Necessary changes in the configuration of the multiplexer can be made in the configuration file [config.go](https://github.com/claygod/door/blob/master/config.go)

# API

Methods:
-  *New* - create a new multiplexer
-  *Add* - add a rule specifying the handler (the default method - GET, ID - as a string to this rule)
-  *Start* - start the server indicating the listening port
-  *Test* - Start analogue (for testing only)

Example:
`
	m := door.New()
	m.Add("/", IHandler)
`

# Named parameters

Arguments in the rules designated route colon. Example route: */abc/:param* , where *abc* is a static section and *:param* - the dynamic section(argument).
The parameters are transmitted via `context`

# Perfomance

- BenchmarkDoorRoot-4     	20000000	       101 ns/op
- BenchmarkDoorTwoSec-4   	 5000000	       309 ns/op
- BenchmarkDoorTenSec-4   	 1000000	      1231 ns/op

# Static files

The directory path to the file and its nickname as part of URL specified in the configuration file. This constants *FILE_PREF* and *FILE_PATH*
