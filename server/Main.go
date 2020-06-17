package main

import (
	"fmt"
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"

	app "github.com/kliver98/api_for_domains"
	"github.com/kliver98/api_for_domains/database"

)

var API_BASE string = "/api/v1/"
var GET_SERVERS string = API_BASE+"domain=:domain"
var GET_SERVERS_SEARCHED string = API_BASE+"history"

func initializeRepo(trc *zipkin.Tracer) (app.DomainDB,app.HistoryDB) {
	var domain app.DomainDB
	domain = newCockroachDomainRepository(trc)
	var history app.DomainDB
	history = newCockroachHistoryRepository(trc)
	return domain, history
}

func newCockroachDomainRepository(trc *zipkin.Tracer, db *sql.DB) app.DomainRepository {

	return database.NewDomainRepository(db, trc)
}

func newCockroachHistoryRepository(trc *zipkin.Tracer, db *sql.DB) app.HistoryRepository {

	return database.NewHistoryRepository(db, trc)
}

func main() { 
	router := fasthttprouter.New()
	router.GET("/", Index)
	router.GET(GET_SERVERS, GetServers)
	router.GET(GET_SERVERS_SEARCHED, GetServersSearched)

	fmt.Println("server starting on localhost:8082")
	log.Fatal(fasthttp.ListenAndServe(":8082", router.Handler))
}