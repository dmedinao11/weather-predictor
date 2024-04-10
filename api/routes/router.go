package routes

import (
	"database/sql"
	"github.com/dmedinao11/weather-predictor/api/handler"
	"github.com/dmedinao11/weather-predictor/internal/weather"
	"github.com/gin-gonic/gin"
)

type (
	Router struct {
		app         *gin.Engine
		db          *sql.DB
		routerGroup *gin.RouterGroup
	}
)

func NewRouter(app *gin.Engine, db *sql.DB) *Router {
	return &Router{app: app, db: db}
}

func (r *Router) MapRoutes() {
	pingHandler := handler.Ping{}
	r.app.GET("/ping", pingHandler.Pong())

	r.routerGroup = r.app.Group("api/v1")
	r.buildWeatherRoutes()
}

func (r *Router) buildWeatherRoutes() {
	repository := weather.NewRepository(r.db)
	service := weather.NewService(&repository)
	weatherHandler := handler.NewWeatherHandler(service)

	r.routerGroup.POST("weather/prediction", weatherHandler.ProcessPrediction())
	r.routerGroup.GET("weather/prediction", weatherHandler.GetPredictionSummary())
	r.routerGroup.GET("weather/prediction/day/:day", weatherHandler.GetPredictionForADay())
}
