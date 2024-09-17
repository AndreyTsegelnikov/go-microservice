package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"context"
	"log"
	"github.com/AndreyTsegelnikov/go-microservice/internal/middleware"
	"github.com/AndreyTsegelnikov/go-microservice/internal/handler"
)

type server struct {
	Router *gin.Engine
	Server *http.Server
}

type Application struct {
	public  server
	private server
	onterm  func()
	stop    chan struct{}
	name    string
	version string
}

// NewApp .
func NewApp(debug bool, name, version string, public, private string) (*Application, chan struct{}) {
	switch debug {
	case true:
		gin.SetMode(gin.DebugMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	app := Application{
		public:  server{Server: &http.Server{Addr: public}, Router: gin.New()},
		private: server{Server: &http.Server{Addr: private}, Router: gin.New()},
		name:    name,
		version: version,
		onterm:  nil,
		stop:    make(chan struct{}),
	}

	app.public.Router.Use(gin.Recovery())
	app.private.Router.Use(gin.Recovery())

	if debug {
		app.public.Router.Use(gin.Logger())
		app.private.Router.Use(gin.Logger())
	}

	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
		<-signals

		if debug {
			log.Println("app terminating...")
		}

		err := app.public.Server.Shutdown(context.Background())
		if err != nil {
			log.Printf("failed to shutdown public server, err: %s\n", err.Error())
		}

		err = app.private.Server.Shutdown(context.Background())
		if err != nil {
			log.Printf("failed to shutdown private server, err: %s\n", err.Error())
		}

		if app.onterm != nil {
			app.onterm()
		}

		err = app.public.Server.Shutdown(context.TODO())
		if err != nil {
			log.Printf("failed to shutdown public server, err: %s\n", err.Error())
		}

		if debug {
			log.Println("app terminated")
		}

		app.stop <- struct{}{}
	}()

	app.public.Router.Use(middleware.Version(version))
	app.private.Router.Use(middleware.Version(version))

	app.private.Router.GET("/liveness", handler.Dummy)
	app.private.Router.GET("/readiness", handler.Readiness)

	app.public.Server.Handler = app.public.Router
	app.private.Server.Handler = app.private.Router

	return &app, app.stop
}

// ServePrivateHTTP starts the app
func (app *Application) ServePrivateHTTP() {
	go func() {
		log.Println(app.private.Server.ListenAndServe())
	}()
}

// ServePublicHTTP starts the app
func (app *Application) ServePublicHTTP() {
	go func() {
		log.Println(app.public.Server.ListenAndServe())
	}()
}

// PublicRouter .
func (app *Application) PublicRouter() *gin.Engine {
	return app.public.Router
}

// PrivateRouter .
func (app *Application) PrivateRouter() *gin.Engine {
	return app.private.Router
}

// OnTerm .
func (app *Application) OnTerm(onterm func()) {
	app.onterm = onterm
}
