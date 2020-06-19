package main

import (
	"fmt"
	"database/sql"
	"strconv"
	"log"
	"time"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"

	database "./database"

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
	main := &Main{}
	router := fasthttprouter.New()
	router.GET("/", Index)
	router.GET(GET_DOMAIN, main.getDomain)
	router.GET(GET_HISTORY, main.getHistory)

	db,_ := database.GetConnection()
	main.router = router
	main.db = db

	return main
}

func main() { 
	main := Init()
	loc, _ := time.LoadLocation("UTC")
	fmt.Println("server starting on localhost:"+strconv.Itoa(PORT)+" >> "+time.Now().In(loc).Format("2006-01-02 15:04:05"))
	log.Fatal(fasthttp.ListenAndServe(":"+strconv.Itoa(PORT), main.Router().Handler))
}