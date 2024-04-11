package main

import (
	"github.com/dmedinao11/weather-predictor/api/routes"
	"github.com/dmedinao11/weather-predictor/pkg/db"
	"github.com/gin-gonic/gin"
	"log"
)

// @title			Weather predictor API 🌦️
// @version		1.0
// @description	API to consult weather predictions
//
// @host			localhost:8080
// @BasePath		/api/v1
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
