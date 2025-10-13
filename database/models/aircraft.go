package models

import (
	"air-trail-backend/database"
	"log"
	"time"
)

// 'icao24','timestamp','acars','adsb','built','categoryDescription','country','engines','firstFlightDate','firstSeen','icaoAircraftClass','lineNumber','manufacturerIcao','manufacturerName','model','modes','nextReg','notes','operator','operatorCallsign','operatorIata','operatorIcao','owner','prevReg','regUntil','registered','registration','selCal','serialNumber','status','typecode','vdl'

// '8a003c','2017-09-15 00:10:03',0,0,,'',Indonesia,'',,,L2J,'',AIRBUS,Airbus,A330 341,0,'','','',INDONESIA,'',GIA,Garuda Indonesia,'',,,PK-GPG,'','165','',A333,0

type Aircraft struct {
	Icao24           string     `gorm:"primaryKey;size:12" json:"icao24"`
	Timestamp        *time.Time `json:"timestamp"`
	Registration     *string    `json:"registration"`
	Country          *string    `json:"country"`
	TypeCode         *string    `json:"type_code"`
	Owner            *string    `json:"owner"`
	ManufacturerIcao *string    `json:"manufacturer_icao"`
	ManufacturerName *string    `json:"manufacturer_name"`
	Model            *string    `json:"model"`
	Operator         *string    `json:"operator"`
	OperatorIcao     *string    `json:"operator_icao"`
	OperatorCallsign *string    `json:"operator_callsign"`
	OperatorIata     *string    `json:"operator_iata"`
}

func (Aircraft) TableName() string {
	return "aircrafts_data"
}

func init() {
	log.Println("[DATABASE] Running auto migration for table aircrafts_data ...")
	if err := database.Pgsql.AutoMigrate(&Aircraft{}); err != nil {
		log.Printf("[DATABASE] Auto migration failed for table aircrafts_data: %v\n", err)
	}
}
