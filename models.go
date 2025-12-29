package main

type APIResponse struct {
	QueryCost       int     `json:"queryCost"`
	Latitude        float64 `json:"latitude"`
	Longtituse      float64 `json:"longtitude"`
	ResolvedAddress string  `json:"resolvedAddress"`
	Address         string  `json:"address"`
	Timezone        string  `json:"timezone"`
	Tzoffset        float64 `json:"tzoffset"`
	Days            []struct {
		DateTime   string   `json:"datetime"`
		Tempmax    float64  `json:"tempmax"`
		Tempmin    float64  `json:"tempmin"`
		Temp       float64  `json:"temp"`
		Humidity   float64  `json:"humidity"`
		Precip     float64  `json:"precip"`
		Windspeed  float64  `json:"windspeed"`
		Pressure   float64  `json:"pressure"`
		CloudCover float64  `json:"cloudcover"`
		Sunrise    string   `json:"sunrise"`
		Conditions string   `json:"conditions"`
		Icon       string   `json:"icon"`
		Stations   []string `json:"stations"`
		Hours      []struct {
			DateTime   string  `json:"datetime"`
			Temp       float64 `json:"temp"`
			Windspeed  float64 `json:"windspeed"`
			Pressure   float64 `json:"pressure"`
			Conditions string  `json:"conditions"`
			CloudCover float64 `json:"cloudcover"`
			Icon       string  `json:"icon"`
			Dew        float64 `json:"dew"`
		} `json:"hours"`
	} `json:"days"`
}
