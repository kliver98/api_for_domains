package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"

	service "github.com/kliver98/api_for_domains/server/service"

)

var API_BASE string = "/api/v1/"
var GET_SERVERS string = API_BASE+"domain=:domain"
var GET_SERVERS_SEARCHED string = API_BASE+"history"

func initializeServices(db *sql.DB) (*service.DomainService, *service.HistoryService) {
	var domain service.DomainService
	domain = newCockroachDomainService(db)
	var history service.HistoryService
	history = newCockroachHistoryService(db)
	return domain, history
}

func newCockroachDomainService(db *sql.DB) service.DomainService {

	return service.NewDomainService(db)
}

func newCockroachHistoryService(db *sql.DB) service.HistoryService {

	return service.NewHistoryService(db)
}

func main() { 
	router := fasthttprouter.New()
	router.GET("/", Index)
	router.GET(GET_SERVERS, GetServers)
	router.GET(GET_SERVERS_SEARCHED, GetServersSearched)

	fmt.Println("server starting on localhost:8082")
	log.Fatal(fasthttp.ListenAndServe(":8082", router.Handler))
}