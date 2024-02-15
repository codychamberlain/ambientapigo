package ambientapigo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const baseURL string = "https://api.ambientweather.net/v1/devices"

// Device struct is a nested struct that maps to Ambient API JSON structure
type Device []struct {
	MacAddress string `json:"macAddress"`
	LastData   struct {
		Dateutc        int64     `json:"dateutc"`
		Winddir        int       `json:"winddir"`
		Windspeedmph   float64   `json:"windspeedmph"`
		Windgustmph    float64   `json:"windgustmph"`
		Maxdailygust   float64   `json:"maxdailygust"`
		Tempf          float64   `json:"tempf"`
		Hourlyrainin   float64   `json:"hourlyrainin"`
		Dailyrainin    float64   `json:"dailyrainin"`
		Weeklyrainin   float64   `json:"weeklyrainin"`
		Monthlyrainin  float64   `json:"monthlyrainin"`
		Totalrainin    float64   `json:"totalrainin"`
		Baromrelin     float64   `json:"baromrelin"`
		Baromabsin     float64   `json:"baromabsin"`
		Humidity       float64   `json:"humidity"`
		Tempinf        float64   `json:"tempinf"`
		Humidityin     float64   `json:"humidityin"`
		Uv             float64   `json:"uv"`
		Solarradiation float64   `json:"solarradiation"`
		FeelsLike      float64   `json:"feelsLike"`
		DewPoint       float64   `json:"dewPoint"`
		FeelsLikein    float64   `json:"feelsLikein"`
		DewPointin     float64   `json:"dewPointin"`
		LastRain       time.Time `json:"lastRain"`
		Tz             string    `json:"tz"`
		Date           time.Time `json:"date"`
	} `json:"lastData"`
	Info struct {
		Name     string `json:"name"`
		Location string `json:"location"`
		Coords   struct {
			Coords struct {
				Lat float64 `json:"lat"`
				Lon float64 `json:"lon"`
			} `json:"coords"`
			Address   string  `json:"address"`
			Location  string  `json:"location"`
			Elevation float64 `json:"elevation"`
			Geo       struct {
				Type        string    `json:"type"`
				Coordinates []float64 `json:"coordinates"`
			} `json:"geo"`
		} `json:"coords"`
	} `json:"info"`
}

// DeviceData is the struct containing data queried by device, based on MAC address
type DeviceData []struct {
	Dateutc        int64     `json:"dateutc"`
	Winddir        int       `json:"winddir"`
	Windspeedmph   float64   `json:"windspeedmph"`
	Windgustmph    float64   `json:"windgustmph"`
	Maxdailygust   float64   `json:"maxdailygust"`
	Tempf          float64   `json:"tempf"`
	Hourlyrainin   float64   `json:"hourlyrainin"`
	Dailyrainin    float64   `json:"dailyrainin"`
	Weeklyrainin   float64   `json:"weeklyrainin"`
	Monthlyrainin  float64   `json:"monthlyrainin"`
	Totalrainin    float64   `json:"totalrainin"`
	Baromrelin     float64   `json:"baromrelin"`
	Baromabsin     float64   `json:"baromabsin"`
	Humidity       float64   `json:"humidity"`
	Tempinf        float64   `json:"tempinf"`
	Humidityin     float64   `json:"humidityin"`
	Uv             float64   `json:"uv"`
	Solarradiation float64   `json:"solarradiation"`
	FeelsLike      float64   `json:"feelsLike"`
	DewPoint       float64   `json:"dewPoint"`
	FeelsLikein    float64   `json:"feelsLikein"`
	DewPointin     float64   `json:"dewPointin"`
	LastRain       time.Time `json:"lastRain"`
	Loc            string    `json:"loc"`
	Date           time.Time `json:"date"`
}

func GetDevices(applicationKey string, apiKey string) Device {
	url := fmt.Sprintf("%s?applicationKey=%s&apiKey=%s", baseURL, applicationKey, apiKey)
	//Go pull the API requests
	response, err := http.Get(url)
	d := Device{}
	if err != nil {
		fmt.Printf("HTTP Request failed due to error %s\n", err)
	} else { //No errors, lets read the request
		if response.StatusCode != 200 {
			fmt.Print(response.Status)
		} else {
			body, _ := io.ReadAll(response.Body)
			err := json.Unmarshal(body, &d)
			if err != nil {
				fmt.Print(err)
			}
		}
	}
	return d
}

func GetDeviceData(mac string, limit int, applicationKey string, apiKey string) DeviceData {
	url := fmt.Sprintf("%s/%s?applicationKey=%s&apiKey=%s&limit=%d", baseURL, mac, applicationKey, apiKey, limit)
	//Go pull the API requests
	response, err := http.Get(url)
	dd := DeviceData{}
	if err != nil {
		fmt.Printf("HTTP Request failed due to error %s\n", err)
	} else { //No errors, lets read the request
		if response.StatusCode != 200 {
			fmt.Print(response.Status)
		} else {
			body, _ := io.ReadAll(response.Body)
			err := json.Unmarshal(body, &dd)
			if err != nil {
				fmt.Print(err)
			}
		}
	}
	return dd
}
