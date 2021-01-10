# Ambient API for Go
This package provides struct types and functions to assist in querying data from the Ambient Weather API.

## Ambient API Documentation
This helper is developed to v1 of the API. For more information see: https://ambientweather.docs.apiary.io/ and https://help.ambientweather.net/help/api/.

## Prereqs
To access the API this helper requres an API Key and Application key both of which can be requested from your account page at https://dashboard.ambientweather.net/account. 

## Installation
To add to your project:  
```go get github.com/codychamberlain/ambientapigo```

Import to your project:  
```Go
import "github.com/codychamberlain/ambientapigo"
```

## Usage
Example application which calls GetDevices and returns a slice of struct Devices. For each device we poll GetDeviceData which is returns a slice of struct DeviceData. Output is entire struct. Access individual data fields via . notation aligned to data struct below.  
```Go
package main

import (
	"fmt"
	"time"

	"github.com/codychamberlain/ambientapigo"
)

func main() {
	applicationKey := ""
	apiKey := ""
	limit := 1 //Number of Device data items to poll. Default is 288. 

	devices := ambientapigo.GetDevices(applicationKey, apiKey)
	for _, d := range devices {
		fmt.Print(d)

		time.Sleep(1000 * time.Millisecond) //Added this due to issues w/ multiple API requests under a second
		data := ambientapigo.GetDeviceData(d.MacAddress, limit, applicationKey, apiKey)
		for _, dd := range data {
			fmt.Print(dd)
		}
	}
}
```
## Known Issues and Quirks
* At this time this has only been tested with 1 type of weather device. YMMV regarding data fields. Please update struct type as necessary if you identify additional fields. 
* Only point in time data is provided. Real-time data via WebSockets is under development
* GetDeviceData doesn't have paramter for endDate. Default is most recent - this is will be included in next commit. Default behavior pulls the most recent data item. 

## Data Types
```Go
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

//DeviceData is the struct containing data queried by device, based on MAC address
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
```
