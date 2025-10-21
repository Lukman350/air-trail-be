package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
)

type Waypoint struct {
	Name      string  `json:"name"`
	Type      string  `json:"type"`
	RouteType string  `json:"route_type"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func Waypoint_Get() ([]Waypoint, error) {
	response, err := http.Get(BASE_URL_NODE + "/geojson/waypoint")

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	fc, err := geojson.UnmarshalFeatureCollection(body)

	if err != nil {
		return nil, err
	}

	var result []Waypoint

	for _, feature := range fc.Features {
		var coordinates [2]float64

		switch geometry := feature.Geometry.(type) {
		case orb.Point:
			coordinates = geometry
		}

		var waypoint Waypoint
		raw, _ := json.Marshal(feature.Properties)
		json.Unmarshal(raw, &waypoint)
		waypoint.Latitude = coordinates[1]
		waypoint.Longitude = coordinates[0]

		result = append(result, waypoint)
	}

	return result, nil
}
