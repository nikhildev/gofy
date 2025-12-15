package models

type WeatherData struct {
	Location    string
	Country     string
	CountryCode string
	Latitude    float64
	Longitude   float64
	Current     CurrentWeather
	Daily       DailyForecast
	Hourly      HourlyForecast
	Timezone    string
	Error       string
}

type CurrentWeather struct {
	Temperature float64
	Condition   string
	Time        string
}

type DailyForecast struct {
	Time          []string
	TempMax       []float64
	TempMin       []float64
	Sunrise       []string
	Sunset        []string
	Precipitation []float64
	WeatherCode   []int
}

type HourlyForecast struct {
	Time          []string
	Temperature   []float64
	Precipitation []float64
	WeatherCode   []int
}

type GeocodingResult struct {
	Results []struct {
		Name        string  `json:"name"`
		Latitude    float64 `json:"latitude"`
		Longitude   float64 `json:"longitude"`
		Country     string  `json:"country"`
		CountryCode string  `json:"country_code"`
		Admin1      string  `json:"admin1"`
	} `json:"results"`
}

type OpenMeteoResponse struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone  string  `json:"timezone"`
	Current   struct {
		Time        string  `json:"time"`
		Temperature float64 `json:"temperature_2m"`
		WeatherCode int     `json:"weather_code"`
	} `json:"current"`
	Daily struct {
		Time          []string  `json:"time"`
		TempMax       []float64 `json:"temperature_2m_max"`
		TempMin       []float64 `json:"temperature_2m_min"`
		Sunrise       []string  `json:"sunrise"`
		Sunset        []string  `json:"sunset"`
		Precipitation []float64 `json:"precipitation_sum"`
		WeatherCode   []int     `json:"weather_code"`
	} `json:"daily"`
	Hourly struct {
		Time          []string  `json:"time"`
		Temperature   []float64 `json:"temperature_2m"`
		Precipitation []float64 `json:"precipitation"`
		WeatherCode   []int     `json:"weather_code"`
	} `json:"hourly"`
}
