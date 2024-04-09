package weather

import (
	"database/sql"
	"fmt"
	"github.com/dmedinao11/weather-predictor/internal/apperrrors"
	"github.com/dmedinao11/weather-predictor/internal/weather/entities"
	"log"
	"strings"
)

type (
	RepositoryImplementation struct {
		db *sql.DB
	}
)

func NewRepository(db *sql.DB) RepositoryImplementation {
	return RepositoryImplementation{
		db: db,
	}
}

func (r *RepositoryImplementation) SaveAll(weatherItems []entities.WeatherItem) error {
	var valueStrings []string
	var valueArgs []interface{}
	for _, item := range weatherItems {
		valueStrings = append(valueStrings, "(?, ?, ?, ?)")
		valueArgs = append(valueArgs, item.Day)
		valueArgs = append(valueArgs, item.WeatherStatus)
		valueArgs = append(valueArgs, item.Perimeter)
		valueArgs = append(valueArgs, item.Period)
	}
	smt := `INSERT INTO weather_item(day_id, weather_status, perimeter, period_id) VALUES %s`
	smt = fmt.Sprintf(smt, strings.Join(valueStrings, ","))

	_, err := r.db.Exec(smt, valueArgs...)

	return err
}

func (r *RepositoryImplementation) CountPeriodsByWeatherStatus(status entities.Status) (int, error) {
	smt := `SELECT COUNT(DISTINCT period_id) FROM weather_item WHERE weather_status=?`
	rows, err := r.db.Query(smt, status)
	if err != nil {
		return 0, err
	}

	defer rows.Close()

	for rows.Next() {
		var count int
		err = rows.Scan(&count)
		return count, err
	}

	return 0, apperrrors.ErrNotFound
}

func (r *RepositoryImplementation) GetAllPeriodsDetail() ([]entities.PeriodDetail, error) {
	smt := `
		SELECT wi.period_id,
			   ws.status,
			   max_rainy.day_id,
			   COUNT(wi.day_id)
		FROM weather_item wi
				 JOIN weather_status ws ON ws.id = wi.weather_status
				 LEFT JOIN (SELECT day_id,
								   period_id
							FROM weather_item
							WHERE (period_id,
								   perimeter) IN (SELECT period_id,
														 MAX(perimeter) AS perimeter
												  FROM weather_item
												  GROUP BY period_id)
							  AND weather_status = 2) max_rainy ON wi.period_id = max_rainy.period_id
		GROUP BY wi.period_id, ws.status, max_rainy.day_id
		ORDER BY wi.period_id
	`

	rows, err := r.db.Query(smt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var periods []entities.PeriodDetail
	for rows.Next() {
		var period entities.PeriodDetail
		if err = rows.Scan(&period.Id, &period.Type, &period.MaxDay, &period.Duration); err != nil {
			log.Fatalln(err.Error())
			return nil, apperrrors.ErrScanFailed
		}
		periods = append(periods, period)
	}

	return periods, nil
}

func (r *RepositoryImplementation) GetPredictionForADay(day int) (entities.WeatherItem, error) {
	prediction := entities.WeatherItem{}
	smt := `SELECT day_id, weather_status, period_id, perimeter FROM weather_item WHERE day_id=?`
	rows, err := r.db.Query(smt, day)
	if err != nil {
		return prediction, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&prediction.Day, &prediction.WeatherStatus, &prediction.Period, &prediction.Perimeter)
		return prediction, err
	}

	return prediction, apperrrors.ErrNotFound
}

func (r *RepositoryImplementation) GetPeakRainIntensityDayForAPeriod(period int) (int, error) {
	smt := `
		SELECT day_id
		FROM weather_item
		WHERE (period_id,
			   perimeter) IN (SELECT period_id,
									 MAX(perimeter) AS perimeter
							  FROM weather_item
							  GROUP BY period_id)
		  AND period_id = ?
	`
	rows, err := r.db.Query(smt, period)
	if err != nil {
		return 0, err
	}

	defer rows.Close()

	for rows.Next() {
		var day int
		err = rows.Scan(&day)
		return day, err
	}

	return 0, apperrrors.ErrNotFound
}
