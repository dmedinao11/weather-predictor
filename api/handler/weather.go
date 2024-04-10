package handler

import (
	"github.com/dmedinao11/weather-predictor/internal/apperrrors"
	"github.com/dmedinao11/weather-predictor/internal/weather/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	dayParamKey = "day"
)

type (
	Service interface {
		ProcessPrediction() error
		GetPredictionSummary() (entities.Prediction, error)
		GetPredictionDetails() ([]entities.PeriodDetail, error)
		GetPredictionForADay(day uint) (entities.WeatherItem, error)
		GetPeakRainIntensityDayForAPeriod(period uint) (int, error)
	}

	Weather struct {
		service Service
	}

	predictionDTO struct {
		DroughtPeriods        int               `json:"drought_periods"`
		RainyPeriods          int               `json:"rainy_periods"`
		OptimalWeatherPeriods int               `json:"optimal_weather_periods"`
		PeriodsDetail         []periodDetailDTO `json:"periods_detail,omitempty"`
		NormalPeriods         int
	}

	periodDetailDTO struct {
		Type     string `json:"type"`
		Duration int    `json:"duration"`
		MaxDay   *int   `json:"max_day,omitempty"`
		Id       int    `json:"id"`
	}

	weatherItemDTO struct {
		Day           int    `json:"day"`
		WeatherStatus string `json:"weather_status"`
		MaxDay        int    `json:"max_day,omitempty"`
		IsPeakDay     bool   `json:"is_peak_day,omitempty"`
	}
)

func NewWeatherHandler(service Service) *Weather {
	return &Weather{service: service}
}

func (w Weather) ProcessPrediction() gin.HandlerFunc {
	return func(context *gin.Context) {
		err := w.service.ProcessPrediction()
		if err == nil {
			context.Status(http.StatusAccepted)
			return
		}
		handleError(context, err)
	}
}

func (w Weather) GetPredictionSummary() gin.HandlerFunc {
	return func(context *gin.Context) {
		summary, err := w.service.GetPredictionSummary()
		if err != nil {
			handleError(context, err)
			return
		}

		var predictionDetails []entities.PeriodDetail
		if context.Query("detailed") == "true" {
			predictionDetails, err = w.service.GetPredictionDetails()
			if err != nil {
				handleError(context, err)
				return
			}
		}

		dto := mapToDto(summary, predictionDetails)
		context.JSON(http.StatusOK, &dto)
	}
}

func (w Weather) GetPredictionForADay() gin.HandlerFunc {
	return func(context *gin.Context) {
		dayParam := context.Param(dayParamKey)
		day, err := strconv.Atoi(dayParam)
		if err != nil {
			handleError(context, apperrrors.ErrParsingPathParam)
			return
		}

		if day < 0 {
			handleError(context, apperrrors.ErrInvalidPathParam)
			return
		}

		predictionForADay, err := w.service.GetPredictionForADay(uint(day))
		if err != nil {
			handleError(context, err)
			return
		}

		dto := mapWeatherItemToDto(predictionForADay)

		if predictionForADay.WeatherStatus == entities.Rainy {
			maxDay, err := w.service.GetPeakRainIntensityDayForAPeriod(uint(predictionForADay.Period))
			dto.IsPeakDay = maxDay == day
			if err != nil {
				handleError(context, err)
				return
			}
		}

		context.JSON(http.StatusOK, &dto)
	}
}

func mapWeatherItemToDto(item entities.WeatherItem) weatherItemDTO {
	return weatherItemDTO{
		Day:           item.Day,
		WeatherStatus: item.WeatherStatus.String(),
	}
}

func mapToDto(summary entities.Prediction, details []entities.PeriodDetail) predictionDTO {
	detailDTOS := make([]periodDetailDTO, 0, len(details))
	for _, detail := range details {
		dto := periodDetailDTO{Id: detail.Id, Type: detail.Type, Duration: detail.Duration, MaxDay: detail.MaxDay}
		detailDTOS = append(detailDTOS, dto)
	}

	return predictionDTO{
		DroughtPeriods:        summary.DroughtPeriods,
		RainyPeriods:          summary.RainyPeriods,
		OptimalWeatherPeriods: summary.OptimalWeatherPeriods,
		NormalPeriods:         summary.NormalPeriods,
		PeriodsDetail:         detailDTOS,
	}
}
