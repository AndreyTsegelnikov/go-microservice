package main

import (
	"go-microservice/internal/app"
	"go-microservice/internal/handler"

	"time"

	"log"
)

var (
	appname = "go-microservice"
	version = "1.0.0"
	build   = "20240915"
	public  = "0.0.0.0:8080"
	private = "0.0.0.0:8081"
	debug   = true
)

func init() {
	log.Println("Started:", time.Now())
	log.Println("App version:", version)
	log.Println("App build:", build)
	log.Println("App name:", appname)
	log.Println(`Private http at:`, `http://`+private)
	log.Println(`Public http at:`, `http://`+public)
}

func main() {
	app, wait := app.NewApp(debug, appname, version, public, private)
	api := app.PublicRouter().Group(appname + "/api/:ver")
	api.GET("/time", handler.Time)
	// старт http
	app.ServePrivateHTTP()
	// старт http
	app.ServePublicHTTP()
	// стоп канал
	<-wait
}
