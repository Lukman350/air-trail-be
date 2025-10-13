package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
)

var Cat021Channel chan Cat021 = make(chan Cat021, 12)
var Cat021Cache sync.Map

type Cat021 struct {
	AsterixType     *string     `json:"asterixType"`
	Sac             *int        `json:"sac"`
	Sic             *int        `json:"sic"`
	SacSicName      *string     `json:"sacSicName"`
	SystemTrackID   *string     `json:"systemTrackID"`
	TimeOfMessage   *float32    `json:"timeOfMessage"`
	FlightLevel     *float32    `json:"flightLevel"`
	Callsign        *string     `json:"callsign"`
	IcaoAddress     *string     `json:"icaoAddress"`
	Speed           *float32    `json:"speed"`
	Heading         *float32    `json:"heading"`
	FirIcao         *string     `json:"firIcao"`
	FirName         *string     `json:"firName"`
	FpDep           *string     `json:"fpDep"`
	FpDest          *string     `json:"fpDest"`
	FpRoute         *string     `json:"fpRoute"`
	AircraftType    *string     `json:"aircraftType"`
	Registration    *string     `json:"registration"`
	UpdateTimestamp *time.Time  `json:"updateTimestamp"`
	UpdateDelete    *string     `json:"updateDelete"`
	Coordinates     *[2]float64 `json:"coordinates"`
	Latitude        float64     `json:"latitude"`
	Longitude       float64     `json:"longitude"`
}

func (data *Cat021) Get() error {
	response, err := http.Get(BASE_URL_BDG + "/cat-021-track")

	if err != nil {
		log.Println(err.Error())
		return err
	}

	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	fc, err := geojson.UnmarshalFeatureCollection(body)

	if err != nil {
		return err
	}

	for _, feature := range fc.Features {
		var coordinates [2]float64
		var cat021 Cat021

		switch geometry := feature.Geometry.(type) {
		case orb.Point:
			coordinates = geometry
		}

		raw, _ := json.Marshal(feature.Properties)
		json.Unmarshal(raw, &cat021)
		cat021.Coordinates = &coordinates
		cat021.Latitude = coordinates[1]
		cat021.Longitude = coordinates[0]
		Cat021Channel <- cat021
	}

	return nil
}
