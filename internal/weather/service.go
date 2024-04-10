package weather

import (
	"github.com/dmedinao11/weather-predictor/internal/weather/entities"
	"math"
)

const (
	years       = 10
	daysInAYear = 365

	concurrentProcess = 10
)

var (
	ferengi = planet{
		speed:           -1,
		distanceFromSun: 500,
	}

	vulcano = planet{
		speed:           5,
		distanceFromSun: 1000,
	}

	betazoide = planet{
		speed:           -3,
		distanceFromSun: 2000,
	}
)

type (
	planet struct {
		speed           int
		distanceFromSun int
	}

	Service struct {
		repo Repository
	}

	Repository interface {
		SaveAll(items []entities.WeatherItem) error
		CountPeriodsByWeatherStatus(status entities.Status) (int, error)
		GetAllPeriodsDetail() ([]entities.PeriodDetail, error)
		GetPredictionForADay(day int) (entities.WeatherItem, error)
		GetPeakRainIntensityDayForAPeriod(period int) (int, error)
	}
)

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (service *Service) ProcessPrediction() error {
	daysToPredict := years * daysInAYear
	weatherItems := make([]entities.WeatherItem, daysToPredict)

	periodId := 0
	var currentWeatherStatus entities.Status
	for i := 0; i < daysToPredict; i++ {
		weatherItem := predictDay(i)

		if weatherItem.WeatherStatus != currentWeatherStatus {
			periodId += 1
			currentWeatherStatus = weatherItem.WeatherStatus
		}

		weatherItem.Period = periodId
		weatherItems[i] = weatherItem
	}

	return service.repo.SaveAll(weatherItems)
}

func (service *Service) GetPredictionSummary() (entities.Prediction, error) {
	optimalStatusCount, err := service.repo.CountPeriodsByWeatherStatus(entities.Optimal)
	if err != nil {
		return entities.Prediction{}, err
	}

	rainyStatusCount, err := service.repo.CountPeriodsByWeatherStatus(entities.Rainy)
	if err != nil {
		return entities.Prediction{}, err
	}

	droughtStatusCount, err := service.repo.CountPeriodsByWeatherStatus(entities.Drought)
	if err != nil {
		return entities.Prediction{}, err
	}

	normalStatusCount, err := service.repo.CountPeriodsByWeatherStatus(entities.Normal)
	if err != nil {
		return entities.Prediction{}, err
	}

	prediction := entities.Prediction{
		DroughtPeriods:        droughtStatusCount,
		RainyPeriods:          rainyStatusCount,
		OptimalWeatherPeriods: optimalStatusCount,
		NormalPeriods:         normalStatusCount,
	}

	return prediction, nil
}

func (service *Service) GetPredictionDetails() ([]entities.PeriodDetail, error) {
	return service.repo.GetAllPeriodsDetail()
}

func (service *Service) GetPredictionForADay(day uint) (entities.WeatherItem, error) {
	return service.repo.GetPredictionForADay(int(day))
}

func (service *Service) GetPeakRainIntensityDayForAPeriod(period uint) (int, error) {
	return service.repo.GetPeakRainIntensityDayForAPeriod(int(period))
}

func predictDay(day int) entities.WeatherItem {
	xF, yF := calculatePosition(ferengi, day)
	xV, yV := calculatePosition(vulcano, day)
	xB, yB := calculatePosition(betazoide, day)

	triangleArea := roundFloat(calculateTriangleArea(xF, yF, xV, yV, xB, yB), 1)
	areAligned := triangleArea == float64(0)

	trianglePerimeter := float64(0)

	if !areAligned {
		trianglePerimeter = roundFloat(calculateTrianglePerimeter(xF, yF, xV, yV, xB, yB), 1)
	}

	weatherStatus := entities.Normal

	if areAligned {
		weatherStatus = entities.Optimal

		if areAlignedWithSun(xF, yF, xV, yV) {
			weatherStatus = entities.Drought
		}
	}

	if !areAligned && isSunInside(triangleArea, xF, yF, xV, yV, xB, yB) {
		weatherStatus = entities.Rainy
	}

	return entities.WeatherItem{
		Day:           day,
		WeatherStatus: weatherStatus,
		Perimeter:     trianglePerimeter,
	}
}

func isSunInside(triangleArea float64, x1 float64, y1 float64, x2 float64, y2 float64, x3 float64, y3 float64) bool {
	xSun := float64(0)
	ySun := float64(0)

	a1 := calculateTriangleArea(xSun, ySun, x2, y2, x3, y3)
	a2 := calculateTriangleArea(x1, y1, xSun, ySun, x3, y3)
	a3 := calculateTriangleArea(x1, y1, x2, y2, xSun, ySun)

	return triangleArea == roundFloat(a1+a2+a3, 1)
}

func areAlignedWithSun(x1 float64, y1 float64, x2 float64, y2 float64) bool {
	if x2 == x1 {
		return true
	}

	slope := (y2 - y1) / (x2 - x1)
	slopeOrigin := y2 / x2

	return roundFloat(slope, 1) == roundFloat(slopeOrigin, 1)

}

func calculateTrianglePerimeter(x1 float64, y1 float64, x2 float64, y2 float64, x3 float64, y3 float64) float64 {
	return calculateDistance(x1, y1, x2, y2) +
		calculateDistance(x2, y2, x3, y3) +
		calculateDistance(x3, y3, x1, y1)
}

func calculateDistance(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

func calculateTriangleArea(x1 float64, y1 float64, x2 float64, y2 float64, x3 float64, y3 float64) float64 {
	a := x1 * (y2 - y3)
	b := x2 * (y3 - y1)
	c := x3 * (y1 - y2)
	return math.Abs((a + b + c) / 2)
}

func calculatePosition(p planet, day int) (float64, float64) {
	angleDegrees := p.speed * day
	angleRadians := float64(angleDegrees) * (math.Pi / float64(180))

	x := float64(p.distanceFromSun) * math.Cos(angleRadians)
	y := float64(p.distanceFromSun) * math.Sin(angleRadians)

	return x, y
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
