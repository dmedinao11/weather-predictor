package main

import (
	"fmt"
	"github.com/dmedinao11/weather-predictor/internal/weather"
	"github.com/dmedinao11/weather-predictor/pkg/db"
)

func main() {
	database, err := db.GetConnection()
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}

	repository := weather.NewRepository(database)
	service := weather.NewService(&repository)

	err = service.ProcessPrediction()
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}
}
