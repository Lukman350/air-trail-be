package models

import "time"

// 'icao24','timestamp','acars','adsb','built','categoryDescription','country','engines','firstFlightDate','firstSeen','icaoAircraftClass','lineNumber','manufacturerIcao','manufacturerName','model','modes','nextReg','notes','operator','operatorCallsign','operatorIata','operatorIcao','owner','prevReg','regUntil','registered','registration','selCal','serialNumber','status','typecode','vdl'

// '8a003c','2017-09-15 00:10:03',0,0,,'',Indonesia,'',,,L2J,'',AIRBUS,Airbus,A330 341,0,'','','',INDONESIA,'',GIA,Garuda Indonesia,'',,,PK-GPG,'','165','',A333,0

type Aircraft struct {
	Icao24           string `gorm:"primaryKey"`
	Timestamp        *time.Time
	Registration     string
	Country          string
	TypeCode         string
	Owner            string
	ManufacturerIcao string
	ManufacturerName string
	Model            string
	Operator         string
	OperatorIcao     string
	OperatorCallsign string
	OperatorIata     string
}

func (Aircraft) TableName() string {
	return "aircrafts_data"
}
