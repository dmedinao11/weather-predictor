package main

import (
	"github.com/dmedinao11/weather-predictor/api/routes"
	"github.com/dmedinao11/weather-predictor/pkg/db"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	database, err := db.GetConnection()
	if err != nil {
		return err
	}

	app := gin.Default()
	router := routes.NewRouter(app, database)
	router.MapRoutes()
	return app.Run()
}
