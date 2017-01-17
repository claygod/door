package door

// The multiplexer `Door`
// Router
// Copyright © 2017 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"log"

	"github.com/valyala/fasthttp"
)

// Router Door is a simple and fast HTTP router for Go (HTTP request multiplexer).
type Router struct {
	routes []*route
	index  *index
	// url    string
}

// New - create a new multiplexer
func New() *Router {
	return &Router{}
}

// Add - add a rule specifying the handler (the default method - GET, ID - as a string to this rule)
func (r *Router) Add(url string, handler func(*fasthttp.RequestCtx)) *route {
	if len(url) > HTTP_PATTERN_COUNT {
		panic("URL is too long")
	} else {
		return r.newRoute(url, handler, []byte(HTTP_METHOD_DEFAULT))
	}
}

// StartFast - start the server indicating the listening port
func (r *Router) Start(port string) {
	r.index = newIndex()
	r.index.compile(r.routes)

	if err := fasthttp.ListenAndServe(port, r.ServeHTTP); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}

// Test - Start analogue (for testing only)
func (r *Router) Test() {
	r.index = newIndex()
	r.index.compile(r.routes)
}
