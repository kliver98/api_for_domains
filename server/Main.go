package main

import (
	"fmt"
	"log"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"

	service "./service"
	database "./database"

)

var API_BASE string = "/api/v1/"
var GET_SERVERS string = API_BASE+"domain=:domain"
var GET_SERVERS_SEARCHED string = API_BASE+"history"


func main() { 
	router := fasthttprouter.New()
	router.GET("/", Index)
	router.GET(GET_SERVERS, GetServers)
	router.GET(GET_SERVERS_SEARCHED, GetServersSearched)

	fmt.Println("server starting on localhost:8082")
	//Testing
	db,_ := database.GetConnection()
	d,_ := service.FetchHistory(db)
	//End testing
	log.Fatal(fasthttp.ListenAndServe(":8082", router.Handler))
}