package models

import (
	"air-trail-backend/database"
	"log"
)

type Airport struct {
	ID                   int64   `gorm:"primaryKey" json:"id"`
	Code                 string  `gorm:"size:12" json:"code"`
	Name                 string  `json:"name"`
	Country              string  `gorm:"size:255" json:"country"`
	Longitude            float64 `json:"longitude"`
	Latitude             float64 `json:"latitude"`
	TotalDayArrivals     *int64  `json:"total_day_arrivals"`
	TotalDayDepartures   *int64  `json:"total_day_departures"`
	TotalMonthArrivals   *int64  `json:"total_month_arrivals"`
	TotalMonthDepartures *int64  `json:"total_month_departures"`
	TotalYearArrivals    *int64  `json:"total_year_arrivals"`
	TotalYearDepartures  *int64  `json:"total_year_departures"`
	TotalOnGround        *int64  `json:"total_on_ground"`
	Status               string  `gorm:"size:24" json:"status"`
	City                 string  `gorm:"size:255" json:"city"`
	Volume               string  `gorm:"size:128" json:"volume"`
	PointType            string  `gorm:"size:255" json:"point_type"`
	Designator           string  `gorm:"size:12" json:"designator"`
}

func (Airport) TableName() string {
	return "airports_data"
}

func Airport_GetAll() ([]Airport, error) {
	var result []Airport
	rows := database.Pgsql.Find(&result)

	if rows.Error != nil {
		return nil, rows.Error
	}

	return result, nil
}

func init() {
	log.Println("[DATABASE] Running auto migration for table airports_data ...")
	if err := database.Pgsql.AutoMigrate(&Airport{}); err != nil {
		log.Printf("[DATABASE] Auto migration failed for table airports_data: %v\n", err)
	}
}
