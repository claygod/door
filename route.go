package door

// The multiplexer `Door`
// Route
// Copyright © 2017 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"bytes"

	"github.com/valyala/fasthttp"
)

// The route for URL
type route struct {
	id       []byte // added by the user
	method   []byte
	handler  func(*fasthttp.RequestCtx)
	sections []*section
}

func (r *Router) newRoute(url string, handler func(*fasthttp.RequestCtx), method []byte) *route {
	route := &route{
		[]byte(url),
		method,
		handler,
		[]*section{},
	}
	route.setSections([]byte(url))
	r.routes = append(r.routes, route)
	return route
}

func (r *route) setSections(url []byte) {
	sec := r.parseUrl(url[1:])
	if len(sec) < HTTP_SECTION_COUNT {
		r.sections = sec
	} else {
		panic("Too many parameters!")
	}
}

func (r *route) Method(value string) *route {
	r.method = []byte(value)
	return r
}

func (r *route) Id(value string) *route {
	r.id = []byte(value)
	return r
}

func (r *route) parseUrl(url []byte) []*section {
	var arraySec []*section
	if len(url) == 0 {
		return []*section{}
	}
	result := r.genSplit(url)

	for _, value := range result {
		if bytes.HasPrefix(value, []byte{DELIMITER_COLON}) {
			arraySec = append(arraySec, newSection(value[1:], TYPE_ARG))
		} else {
			arraySec = append(arraySec, newSection(value, TYPE_STAT))
		}
	}
	return arraySec
}

func (r *route) genSplit(s []byte) [][]byte {
	n := 1
	c := DELIMITER_SLASH
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			n++
		}
	}
	out := make([][]byte, n)
	count := 0
	begin := 0
	length := len(s) - 1
	for i := 0; i <= length; i++ {
		if s[i] == c {
			out[count] = s[begin:i]
			count++
			begin = i + 1
		}
	}
	out[count] = s[begin : length+1]
	return out
}
