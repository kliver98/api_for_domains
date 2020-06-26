package main

import (
	"fmt"
	"database/sql"
	"strconv"
	"log"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"

	database "github.com/kliver98/api_for_domains/server/database"

)

var API_BASE string = "/api/v1/"
var GET_DOMAIN string = API_BASE+"domain=:domain"
var GET_HISTORY string = API_BASE+"history"
var PORT = 8082

type Main struct {
	router *fasthttprouter.Router
	db *sql.DB
}

func (main *Main) Router() *fasthttprouter.Router {
	return main.router
}

func Init() *Main {
	db,_ := database.GetConnection()

	//if err!=nil { //If does not connect with database continuing the flow execution have no meaning
	//	panic(err.Error())
	//}

	main := &Main{}
	router := fasthttprouter.New()
	router.GET("/", Index)
	router.GET(GET_DOMAIN, main.getDomain)
	router.GET(GET_HISTORY, main.getHistory)

	main.router = router
	main.db = db

	return main
}

func main() { 
	main := Init()
	fmt.Println("server starting on localhost:"+strconv.Itoa(PORT))
	log.Fatal(fasthttp.ListenAndServe(":"+strconv.Itoa(PORT), main.Router().Handler))
}