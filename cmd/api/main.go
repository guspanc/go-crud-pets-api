package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/guspanc/go-crud-pets-api/cmd/api/data"
	"github.com/guspanc/go-crud-pets-api/pkg/exit"
	"github.com/guspanc/go-crud-pets-api/pkg/logger"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql" // mysql dialect
)

// Application struct
type Application struct {
	db     data.Datastore
	config *configuration
}

var app *Application

func main() {
	// build app
	app = &Application{}
	app.config = getConfig()
	db, err := data.NewDB(fmt.Sprintf("%s:%s@tcp(%s)/pets", app.config.dbUsername, app.config.dbPassword, app.config.dbEndpoint))
	if err != nil {
		panic(err)
	}
	app.db = db

	// build server
	r := mux.NewRouter()
	registerRoutes(r, app)
	srv := &http.Server{
		Handler: r,
		Addr:    app.config.apiPort,
	}

	// start server
	go func() {
		logger.INFO.Println("starting server on", app.config.apiPort)
		if err := srv.ListenAndServe(); err != nil {
			logger.INFO.Fatal(err)
		}
	}()

	// blocking exit handler
	exit.Init(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			logger.INFO.Println("ERROR", err)
		}
		if err := app.db.Close(ctx); err != nil {
			logger.INFO.Println("ERROR", err)
		}
	})
}
