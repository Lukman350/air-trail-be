package api

import (
	"air-trail-backend/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// [
//     {
//         "datetime": "2025-10-06 10:03:51",
//         "registration": "PKGUD",
//         "departure": "WAHS",
//         "destination": "WIII",
//         "duration": "0050",
//         "speed": "N042",
//         "altitude": "F260",
//         "route": "ANY W17N KIDET W16 IMU CKG ",
//         "raw":"NDP5554 071058      FF WRRRYAYD 071058 WAHSZPZX DEP-GIA239-WAHS1057-WIII-DOF/251007)",
//         "dof": "251007",
//         "etaProgress": 17,
//         "message_type": "DEP",
//         "call_sign": "GIA239",
//         "aircraft_type": "B738",
//         "departure_time": "1055",
//         "destination_time": "1057",
//         "schedule_departure_utc": "2025-10-07T10:55:51Z",
//         "schedule_destination_utc": "2025-10-07T11:45:51Z",
//         "actual_departure_utc": "2025-10-07T10:57:00Z",
//         "departure_airport": {
//             "city": "JENDERAL AHMAD YANI",
//             "country": "Indonesia",
//             "airport_name": "JENDERAL AHMAD YANI"
//         },
//         "destination_airport": {
//             "city": "SOEKARNO-HATTA",
//             "country": "Indonesia",
//             "airport_name": "SOEKARNO-HATTA"
//         }
//     }
// ]

type Airport struct {
	City        string `json:"city"`
	Country     string `json:"country"`
	AirportName string `json:"airport_name"`
}

type Aftn struct {
	DateTime               string  `json:"datetime"`
	Registration           string  `json:"registration"`
	Departure              string  `json:"departure"`
	Destination            string  `json:"destination"`
	Duration               string  `json:"duration"`
	Speed                  string  `json:"speed"`
	Altitude               string  `json:"altitude"`
	Route                  string  `json:"route"`
	Raw                    string  `json:"raw"`
	Dof                    string  `json:"dof"`
	EtaProgress            int64   `json:"etaProgress"`
	MessageType            string  `json:"message_type"`
	CallSign               *string `json:"call_sign"`
	AircraftType           string  `json:"aircraft_type"`
	DepartureTime          string  `json:"departure_time"`
	DestinationTime        string  `json:"destination_time"`
	ScheduleDepartureUTC   string  `json:"schedule_departure_utc"`
	ScheduleDestinationUTC string  `json:"schedule_destination_utc"`
	ActualDepartureUTC     string  `json:"actual_departure_utc"`
	DepartureAirport       Airport `json:"departure_airport"`
	DestinationAirport     Airport `json:"destination_airport"`
}

func (aftn *Aftn) GetByCallsign(callsign string) error {
	response, err := http.Get(fmt.Sprintf("%s/get-aftn/%s", BASE_URL, callsign))

	if err != nil {
		return err
	}

	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)
	json.Unmarshal(body, aftn)

	if aftn.CallSign == nil {
		return &utils.NotFoundError{Message: fmt.Sprintf("No aftn data found for %s", callsign)}
	}

	return nil
}
