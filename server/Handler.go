package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"encoding/json"
	service "github.com/kliver98/api_for_domains/server/service"
)

func Index(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Created by Kliver Daniel Girón\n\n	>>>	"+GET_DOMAIN+"\n	>>>	"+GET_HISTORY)
}

func (main *Main) getDomain(ctx *fasthttp.RequestCtx) {
	param := ctx.UserValue("domain")
	domain,_ := service.FetchDomain(main.db, fmt.Sprint(param))
	jsonBody, err := json.Marshal(domain)

	if err != nil {
		ctx.Error(" json marshal fail", 500)
		return
	}

	ctx.Response.SetBody(jsonBody)
}

func (main *Main) getHistory(ctx *fasthttp.RequestCtx) {
	history,_ := service.FetchHistory(main.db)

	jsonBody, err := json.Marshal(history)

	if err != nil {
		ctx.Error(" json marshal fail", 500)
		return
	}

	ctx.Response.SetBody(jsonBody)
}