package entities

const (
	Optimal Status = 1
	Rainy   Status = 2
	Drought Status = 3
	Normal  Status = 4
)

type (
	Status int

	WeatherItem struct {
		Day           int
		WeatherStatus Status
		Period        int
		Perimeter     float64
	}

	Prediction struct {
		DroughtPeriods        int
		RainyPeriods          int
		OptimalWeatherPeriods int
		NormalPeriods         int
	}

	PeriodDetail struct {
		Id       int
		Type     string
		Duration int
		MaxDay   *int
	}
)

func (s Status) String() string {
	switch s {
	case Optimal:
		return "optimal"
	case Rainy:
		return "rainy"
	case Drought:
		return "drought"
	case Normal:
		return "normal"
	}
	return ""
}
