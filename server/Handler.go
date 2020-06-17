package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func Index(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Created by Kliver Daniel Girón\n\n	>>>	"+GET_SERVERS+"\n	>>>	"+GET_SERVERS_SEARCHED)
}

func GetServers(ctx *fasthttp.RequestCtx) {
	domain := ctx.UserValue("domain")
	fmt.Fprintf(ctx, "Searching for domain: , %s!\n", domain)
}

func GetServersSearched(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Searching for history")
}