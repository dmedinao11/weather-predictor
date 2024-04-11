package weather

import (
	"github.com/dmedinao11/weather-predictor/internal/weather/entities"
	"reflect"
	"testing"
)

func Test_predictDay(t *testing.T) {
	tests := []struct {
		name string
		day  int
		want entities.WeatherItem
	}{
		{
			name: "Day 0",
			day:  0,
			want: entities.WeatherItem{
				Day:           0,
				WeatherStatus: entities.Drought,
				Perimeter:     0,
			},
		},
		{
			name: "Day 1",
			day:  1,
			want: entities.WeatherItem{
				Day:           1,
				WeatherStatus: entities.Normal,
				Perimeter:     3025.1,
			},
		},
		{
			name: "Day 90",
			day:  90,
			want: entities.WeatherItem{
				Day:           90,
				WeatherStatus: entities.Optimal,
				Perimeter:     0,
			},
		},
		{
			name: "Day 270",
			day:  270,
			want: entities.WeatherItem{
				Day:           270,
				WeatherStatus: entities.Optimal,
				Perimeter:     0,
			},
		},
		{
			name: "Day 361",
			day:  361,
			want: entities.WeatherItem{
				Day:           361,
				WeatherStatus: entities.Normal,
				Perimeter:     3025.1,
			},
		},
		{
			name: "Day 450",
			day:  450,
			want: entities.WeatherItem{
				Day:           450,
				WeatherStatus: entities.Optimal,
				Perimeter:     0,
			},
		},
		{
			name: "Day 364",
			day:  364,
			want: entities.WeatherItem{
				Day:           364,
				WeatherStatus: entities.Normal,
				Perimeter:     3354.5,
			},
		},
		{
			name: "Day 365",
			day:  365,
			want: entities.WeatherItem{
				Day:           365,
				WeatherStatus: entities.Normal,
				Perimeter:     3521.1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := predictDay(tt.day); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("predictDay() = %v, want %v", got, tt.want)
			}
		})
	}
}
