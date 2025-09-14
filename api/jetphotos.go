package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type JetPhotoImage struct {
	Image        string `json:"Image"`
	Link         string `json:"Link"`
	Thumbnail    string `json:"Thumbnail"`
	DateTaken    string `json:"DateTaken"`
	DateUploaded string `json:"DateUploaded"`
	Location     string `json:"Location"`
	Photographer string `json:"Photographer"`
	Aircraft     string `json:"Aircraft"`
	Serial       string `json:"Serial"`
	Airline      string `json:"Airline"`
}

type JetPhotosResponse struct {
	Reg    string          `json:"Reg"`
	Images []JetPhotoImage `json:"Images"`
}

func GetJetPhotos(registration string) (JetPhotosResponse, error) {
	response, err := http.Get(fmt.Sprintf("%s?reg=%s&only_jp=true", JETPHOTOS_URL, registration))

	if err != nil {
		return JetPhotosResponse{}, err
	}

	defer response.Body.Close()

	result := JetPhotosResponse{}
	body, _ := io.ReadAll(response.Body)
	json.Unmarshal(body, &result)

	return result, nil
}
