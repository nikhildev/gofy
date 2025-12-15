package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/nikhildev/gofy/web/models"
)

var weatherTmpl *template.Template

func init() {
	var err error
	funcMap := template.FuncMap{
		"formatDate": func(dateStr string) string {
			t, err := time.Parse("2006-01-02", dateStr)
			if err != nil {
				return dateStr
			}
			if time.Now().Format("2006-01-02") == dateStr {
				return "Today"
			}
			if time.Now().AddDate(0, 0, 1).Format("2006-01-02") == dateStr {
				return "Tomorrow"
			}
			return t.Format("Mon, Jan 2")
		},
		"formatTime": func(timeStr string) string {
			t, err := time.Parse("2006-01-02T15:04", timeStr)
			if err != nil {
				return timeStr
			}
			return t.Format("3:04 PM")
		},
		"getCondition": func(code int) string {
			return getWeatherCondition(code)
		},
		"getFlag": func(countryCode string) string {
			return getCountryFlag(countryCode)
		},
		"lower": func(s string) string {
			return strings.ToLower(s)
		},
	}
	weatherTmpl, err = template.New("weather.tmpl").Funcs(funcMap).ParseFiles("templates/weather.tmpl")
	if err != nil {
		log.Fatal("Error parsing weather template:", err)
	}
}

func getWeatherCondition(code int) string {
	conditions := map[int]string{
		0:  "Clear sky",
		1:  "Mainly clear",
		2:  "Partly cloudy",
		3:  "Overcast",
		45: "Foggy",
		48: "Foggy",
		51: "Light drizzle",
		53: "Drizzle",
		55: "Heavy drizzle",
		61: "Light rain",
		63: "Rain",
		65: "Heavy rain",
		71: "Light snow",
		73: "Snow",
		75: "Heavy snow",
		77: "Snow grains",
		80: "Light showers",
		81: "Showers",
		82: "Heavy showers",
		85: "Light snow showers",
		86: "Snow showers",
		95: "Thunderstorm",
		96: "Thunderstorm with hail",
		99: "Thunderstorm with hail",
	}
	if condition, ok := conditions[code]; ok {
		return condition
	}
	return "Unknown"
}

func getCountryFlag(countryCode string) string {
	if len(countryCode) != 2 {
		return ""
	}

	// Convert country code to flag emoji using regional indicator symbols
	// A = U+1F1E6, B = U+1F1E7, etc.
	flag := ""
	for _, char := range countryCode {
		if char >= 'A' && char <= 'Z' {
			flag += string(rune(0x1F1E6 + (char - 'A')))
		}
	}
	return flag
}

func fetchWeatherData(lat, lon float64) (*models.OpenMeteoResponse, error) {
	apiURL := fmt.Sprintf(
		"https://api.open-meteo.com/v1/forecast?latitude=%.4f&longitude=%.4f&current=temperature_2m,weather_code&daily=temperature_2m_max,temperature_2m_min,sunrise,sunset,precipitation_sum,weather_code&hourly=temperature_2m,precipitation,weather_code&timezone=auto&forecast_days=10",
		lat, lon,
	)

	resp, err := http.Get(apiURL)
	if err != nil {
		log.Printf("Error fetching weather data: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Weather API returned status: %d", resp.StatusCode)
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var data models.OpenMeteoResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Printf("Error decoding weather data: %v", err)
		return nil, err
	}

	return &data, nil
}

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	location := r.URL.Query().Get("location")
	if location == "" {
		location = "Espoo"
	}

	// Geocode the location
	geocodeURL := fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1&language=en&format=json",
		url.QueryEscape(location))

	resp, err := http.Get(geocodeURL)
	if err != nil {
		log.Printf("Geocoding error for '%s': %v", location, err)
		renderWeatherError(w, "Failed to fetch location data")
		return
	}
	defer resp.Body.Close()

	var geocode models.GeocodingResult
	if err := json.NewDecoder(resp.Body).Decode(&geocode); err != nil {
		log.Printf("Error decoding geocoding response: %v", err)
		renderWeatherError(w, "Failed to parse location data")
		return
	}

	if len(geocode.Results) == 0 {
		log.Printf("No results found for location: %s", location)
		renderWeatherError(w, fmt.Sprintf("Location '%s' not found", location))
		return
	}

	result := geocode.Results[0]
	log.Printf("Fetching weather for %s (%.4f, %.4f)", result.Name, result.Latitude, result.Longitude)

	weatherData, err := fetchWeatherData(result.Latitude, result.Longitude)
	if err != nil {
		log.Printf("Failed to fetch weather data: %v", err)
		renderWeatherError(w, "Failed to fetch weather data")
		return
	}

	countryCode := result.CountryCode
	if countryCode == "" {
		countryCode = result.Country
	}
	log.Printf("Country: %s, CountryCode: %s", result.Country, countryCode)

	data := models.WeatherData{
		Location:    result.Name,
		Country:     result.Country,
		CountryCode: countryCode,
		Latitude:    result.Latitude,
		Longitude:   result.Longitude,
		Timezone:    weatherData.Timezone,
		Current: models.CurrentWeather{
			Temperature: weatherData.Current.Temperature,
			Condition:   getWeatherCondition(weatherData.Current.WeatherCode),
			Time:        weatherData.Current.Time,
		},
		Daily: models.DailyForecast{
			Time:          weatherData.Daily.Time,
			TempMax:       weatherData.Daily.TempMax,
			TempMin:       weatherData.Daily.TempMin,
			Sunrise:       weatherData.Daily.Sunrise,
			Sunset:        weatherData.Daily.Sunset,
			Precipitation: weatherData.Daily.Precipitation,
			WeatherCode:   weatherData.Daily.WeatherCode,
		},
		Hourly: models.HourlyForecast{
			Time:          weatherData.Hourly.Time,
			Temperature:   weatherData.Hourly.Temperature,
			Precipitation: weatherData.Hourly.Precipitation,
			WeatherCode:   weatherData.Hourly.WeatherCode,
		},
	}

	err = weatherTmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error executing weather template:", err)
	}
}

func renderWeatherError(w http.ResponseWriter, errorMsg string) {
	data := models.WeatherData{
		Location: "Espoo",
		Error:    errorMsg,
	}
	weatherTmpl.Execute(w, data)
}

// API endpoint for location search
func SearchLocationHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		json.NewEncoder(w).Encode([]interface{}{})
		return
	}

	geocodeURL := fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=5&language=en&format=json",
		url.QueryEscape(query))

	resp, err := http.Get(geocodeURL)
	if err != nil {
		http.Error(w, "Failed to search locations", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var geocode models.GeocodingResult
	if err := json.NewDecoder(resp.Body).Decode(&geocode); err != nil {
		http.Error(w, "Failed to parse locations", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(geocode.Results)
}
