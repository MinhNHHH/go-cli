package weather

import (
	"encoding/json"
	"io"
	"net/http"
)

type WeatherResponse struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

type Location struct {
	Name      string `json:"name"`
	Region    string `json:"region"`
	Country   string `json:"country"`
	LocalTime string `json:"localtime"`
}

type Condition struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
	Code int    `json:"code"`
}

type Current struct {
	LastUpdatedEpoch int64     `json:"last_updated_epoch"`
	LastUpdated      string    `json:"last_updated"`
	TemperatureC     float64   `json:"temp_c"`
	TemperatureF     float64   `json:"temp_f"`
	IsDay            int       `json:"is_day"`
	Condition        Condition `json:"condition"`
	WindMph          float64   `json:"wind_mph"`
	WindKph          float64   `json:"wind_kph"`
	WindDegree       int       `json:"wind_degree"`
	WindDirection    string    `json:"wind_dir"`
	PressureMb       float64   `json:"pressure_mb"`
	PressureIn       float64   `json:"pressure_in"`
	PrecipitationMm  float64   `json:"precip_mm"`
	PrecipitationIn  float64   `json:"precip_in"`
	Humidity         int       `json:"humidity"`
	Cloud            int       `json:"cloud"`
	FeelsLikeC       float64   `json:"feelslike_c"`
	FeelsLikeF       float64   `json:"feelslike_f"`
	VisibilityKm     float64   `json:"vis_km"`
	VisibilityMiles  float64   `json:"vis_miles"`
	UV               float64   `json:"uv"`
	GustMph          float64   `json:"gust_mph"`
	GustKph          float64   `json:"gust_kph"`
}

func GetWeatherInform() WeatherResponse {
	res, err := http.Get("http://api.weatherapi.com/v1/current.json?key=9b1afd43554842619e372827241803&q=Hanoi")
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather API not available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	// Define a variable to hold the weather response
	var weatherResp WeatherResponse

	// Unmarshal the JSON data into the WeatherResponse struct
	err = json.Unmarshal(body, &weatherResp)
	if err != nil {
		panic(err)
	}

	return weatherResp
}
