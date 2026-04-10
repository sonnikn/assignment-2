package models

// RegistrationRequest representerer innholdet fra en POST/PUT forespørsel[cite: 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85].
type Registration struct {
	ID               string       `json:"id" firestore:"id"`
	Country          string       `json:"country,omitempty" firestore:"country,omitempty"`
	IsoCode          string       `json:"isoCode,omitempty" firestore:"isoCode,omitempty"`
	Features         FeatureFlags `json:"features" firestore:"features"`
	TargetCurrencies []string     `json:"targetCurrencies" firestore:"targetCurrencies"`
	LastChange       string       `json:"lastChange" firestore:"lastChange"` // Format: YYYYMMDD HH:mm
}

type FeatureFlags struct {
	Temperature   bool `json:"temperature" firestore:"temperature"`
	Precipitation bool `json:"precipitation" firestore:"precipitation"`
	AirQuality    bool `json:"airQuality" firestore:"airQuality"`
	Capital       bool `json:"capital" firestore:"capital"`
	Coordinates   bool `json:"coordinates" firestore:"coordinates"`
	Population    bool `json:"population" firestore:"population"`
	Area          bool `json:"area" firestore:"area"`
}

// Dashboard representerer en ferdig utfylt dashboard respons
type Dashboard struct {
	Country       string            `json:"country"`
	IsoCode       string            `json:"isoCode"`
	Features      DashboardFeatures `json:"features"`
	LastRetrieval string            `json:"lastRetrieval"`
}

type DashboardFeatures struct {
	Temperature      *float64           `json:"temperature,omitempty"`
	Precipitation    *float64           `json:"precipitation,omitempty"`
	AirQuality       *AirQualityData    `json:"airQuality,omitempty"`
	Capital          *string            `json:"capital,omitempty"`
	Coordinates      *CoordinatesData   `json:"coordinates,omitempty"`
	Population       *int               `json:"population,omitempty"`
	Area             *float64           `json:"area,omitempty"`
	TargetCurrencies map[string]float64 `json:"targetCurrencies,omitempty"`
}

type AirQualityData struct {
	Pm25  float64 `json:"pm25"`
	Pm10  float64 `json:"pm10"`
	Level string  `json:"level"`
}

type CoordinatesData struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
