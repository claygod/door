package door

// The multiplexer `Door`
// Tests
// Copyright © 2017 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestRouting(t *testing.T) {
	ctx := getCtxForTest("GET /abc/10 HTTP/1.1\nHost: aaa.com\n\n")

	mux := New()
	mux.Add("/abc/:par", func(ctx *fasthttp.RequestCtx) { ctx.SetUserValue("x", "y") }).Method("GET")
	mux.Test()
	mux.ServeHTTP(ctx)

	if ctx.UserValue("x") != "y" {
		t.Error("The expected value of `y`, and received:", ctx.UserValue("x"))
	}
}

func TestError(t *testing.T) {
	ctx := getCtxForTest("GET /qwerty/10 HTTP/1.1\nHost: aaa.com\n\n")

	mux := New()
	mux.Add("/abc/:par", func(ctx *fasthttp.RequestCtx) { ctx.SetUserValue("x", "y") }).Method("GET")
	mux.Test()
	mux.ServeHTTP(ctx)

	if ctx.UserValue("x") != nil {
		t.Error("A router incorrectly chose Route")
	}
}

func TestMethod(t *testing.T) {
	ctx := getCtxForTest("GET /abc/10 HTTP/1.1\nHost: aaa.com\n\n")

	mux := New()
	mux.Add("/abc/:par", func(ctx *fasthttp.RequestCtx) { ctx.SetUserValue("x", "y") }).Method("POST")
	mux.Test()
	mux.ServeHTTP(ctx)

	if ctx.UserValue("x") != nil {
		t.Error("Reversed methods GET and POST")
	}
}

// Test if the mux don't handle by prefix (static)
func TestPathStatic(t *testing.T) {
	ctx := getCtxForTest("GET /a/b HTTP/1.1\nHost: aaa.com\n\n")

	mux := New()
	mux.Add("/a", func(ctx *fasthttp.RequestCtx) { ctx.SetUserValue("x", "1") }).Method("GET")
	mux.Add("/a/b", func(ctx *fasthttp.RequestCtx) { ctx.SetUserValue("x", "2") }).Method("GET")
	mux.Test()
	mux.ServeHTTP(ctx)

	if ctx.UserValue("x") == 1 {
		t.Error("response with the wrong path")
	}
}

// Test if the mux don't handle by prefix (dinamic)
func TestPathDinamic(t *testing.T) {
	ctx := getCtxForTest("GET /abc/10 HTTP/1.1\nHost: aaa.com\n\n")

	mux := New()
	mux.Add("/abc/:par", func(ctx *fasthttp.RequestCtx) { ctx.SetUserValue("x", ctx.UserValue("par")) }).Method("GET")
	mux.Test()
	mux.ServeHTTP(ctx)

	if ctx.UserValue("x") != "10" {
		t.Error("The expected value of `10`, and received:", ctx.UserValue("x"))
	}
}

func TestDefaultMethodGet(t *testing.T) {
	ctx := getCtxForTest("GET /abc/10 HTTP/1.1\nHost: aaa.com\n\n")

	mux := New()
	mux.Add("/abc/:par", func(ctx *fasthttp.RequestCtx) { ctx.SetUserValue("x", "y") })
	mux.Test()
	mux.ServeHTTP(ctx)

	if ctx.UserValue("x") != "y" {
		t.Error("The default must be a GET method")
	}
}

func TestRouteSlash(t *testing.T) {
	ctx := getCtxForTest("GET / HTTP/1.1\nHost: aaa.com\n\n")

	mux := New()
	mux.Add("/", func(ctx *fasthttp.RequestCtx) { ctx.SetUserValue("x", "1") }).Method("GET")
	mux.Add("/abc", func(ctx *fasthttp.RequestCtx) { ctx.SetUserValue("x", "2") }).Method("GET")
	mux.Test()
	mux.ServeHTTP(ctx)

	if ctx.UserValue("x") != "1" {
		t.Error("Error in determining the root `route`")
	}
}

func TestSlashEnd(t *testing.T) {
	ctx := getCtxForTest("GET /abc/ HTTP/1.1\nHost: aaa.com\n\n")

	mux := New()
	mux.Add("/abc", func(ctx *fasthttp.RequestCtx) { ctx.SetUserValue("x", "y") }).Method("GET")
	mux.Test()
	mux.ServeHTTP(ctx)

	if ctx.UserValue("x") == "y" {
		t.Error("Slash removing doesn't work !")
	}
}

func BenchmarkDoorRoot(b *testing.B) {
	ctx := getCtxForTest("GET / HTTP/1.1\nHost: aaa.com\n\n")

	mux := New()
	mux.Add("/", func(ctx *fasthttp.RequestCtx) {}).Method("GET")
	mux.Test()
	for n := 0; n < b.N; n++ {
		mux.ServeHTTP(ctx)
	}
	mux.ServeHTTP(ctx)
}

func BenchmarkDoorTwoSec(b *testing.B) {
	ctx := getCtxForTest("GET /abc/10 HTTP/1.1\nHost: aaa.com\n\n")

	mux := New()
	mux.Add("/abc/:par", func(ctx *fasthttp.RequestCtx) {}).Method("GET")
	mux.Test()
	for n := 0; n < b.N; n++ {
		mux.ServeHTTP(ctx)
	}
	mux.ServeHTTP(ctx)
}

func BenchmarkDoorTenSec(b *testing.B) {
	ctx := getCtxForTest("GET /a/1/b/2/c/3/d/4/e/5 HTTP/1.1\nHost: aaa.com\n\n")

	mux := New()
	mux.Add("/a/:1/b/:2/c/:3/d/:4/e/:5", func(ctx *fasthttp.RequestCtx) {}).Method("GET")
	mux.Test()
	for n := 0; n < b.N; n++ {
		mux.ServeHTTP(ctx)
	}
	mux.ServeHTTP(ctx)
}

func getCtxForTest(str string) *fasthttp.RequestCtx {
	var ctx fasthttp.RequestCtx
	br := bufio.NewReader(bytes.NewBufferString(str))
	if err := ctx.Request.Read(br); err != nil {
		panic(err)
	}
	return &ctx
}
