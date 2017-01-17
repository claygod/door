package door

// The multiplexer `Door`
// Server
// Copyright © 2017 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

// ServeHTTP looks for a matching route
func (r *Router) ServeHTTP(ctx *fasthttp.RequestCtx) {
	if route := r.index.find(ctx.URI().Path(), ctx.Method()); route != nil {
		query := route.genSplit(ctx.URI().Path()[1:])
		for u := len(route.sections) - 1; u >= 0; u-- {
			if route.sections[u].typeSec == TYPE_ARG {
				ctx.SetUserValue(string(route.sections[u].id), string(query[u]))
			}
		}
		route.handler(ctx)
	} else {
		r.Default(ctx)
	}
}

// Default Handler
func (r *Router) Default(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "404 Page not found! Requested path is %q.", ctx.Path())
}
