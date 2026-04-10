package models

import "time"

// ==========================================
// 1. DASHBOARD REGISTRERING (/registrations)
// ==========================================

// RegistrationRequestBody brukes når en bruker oppretter/oppdaterer (POST/PUT)
type RegistrationRequestBody struct {
	Country  string   `json:"country" firestore:"country"`
	IsoCode  string   `json:"isoCode" firestore:"isoCode"`
	Features Features `json:"features" firestore:"features"`
}

// Registration er hvordan dashboardet faktisk lagres i Firestore og returneres via GET
type Registration struct {
	ID            string    `json:"id" firestore:"id"` // Settes av databasen
	Country       string    `json:"country" firestore:"country"`
	IsoCode       string    `json:"isoCode" firestore:"isoCode"`
	Features      Features  `json:"features" firestore:"features"`
	LastRetrieval time.Time `json:"lastRetrieval" firestore:"lastRetrieval"`
}

// Features angir hvilke data brukeren vil at dashboardet skal vise (true/false)
type Features struct {
	Temperature   bool `json:"temperature" firestore:"temperature"`
	Precipitation bool `json:"precipitation" firestore:"precipitation"`
	Capital       bool `json:"capital" firestore:"capital"`
	Coordinates   bool `json:"coordinates" firestore:"coordinates"`
	Population    bool `json:"population" firestore:"population"`
	Area          bool `json:"area" firestore:"area"`
	Borders       bool `json:"borders" firestore:"borders"`
}

// ==========================================
// 2. FERDIG UTFYLT DASHBOARD (/dashboards)
// ==========================================

// Dashboard er svaret du sender ut når noen gjør en GET på /dashboards/{id}
type Dashboard struct {
	ID            string        `json:"id"`
	Country       string        `json:"country"`
	IsoCode       string        `json:"isoCode"`
	Features      DashboardData `json:"features"`
	LastRetrieval time.Time     `json:"lastRetrieval"`
}

// DashboardData inneholder de FAKTISKE verdiene fra API-ene (ikke bools)
// Vi bruker "pointers" (f.eks. *float64) og "omitempty" slik at felter som
// brukeren IKKE har bedt om, skjules helt fra JSON-responsen.
type DashboardData struct {
	Temperature   *float64 `json:"temperature,omitempty"`
	Precipitation *float64 `json:"precipitation,omitempty"`
	Capital       string   `json:"capital,omitempty"`
	Coordinates   *Coords  `json:"coordinates,omitempty"`
	Population    *int64   `json:"population,omitempty"`
	Area          *float64 `json:"area,omitempty"`
	Borders       []string `json:"borders,omitempty"`
}

type Coords struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// ==========================================
// 3. WEBHOOKS (/notifications)
// ==========================================

// WebhookRegistration brukes for inngående data på POST /notifications/
type WebhookRegistration struct {
	ID        string     `json:"id" firestore:"id"`
	URL       string     `json:"url" firestore:"url"`
	Country   string     `json:"country" firestore:"country"`
	Event     string     `json:"event" firestore:"event"`                             // REGISTER, CHANGE, INVOKE, THRESHOLD
	Threshold *Threshold `json:"threshold,omitempty" firestore:"threshold,omitempty"` // Kun for THRESHOLD-events
}

// Threshold definerer en spesifikk regel (f.eks: "temperature", ">", 20.5)
type Threshold struct {
	Field    string  `json:"field" firestore:"field"`
	Operator string  `json:"operator" firestore:"operator"`
	Value    float64 `json:"value" firestore:"value"`
}

// WebhookPayload er den faktiske JSON-strukturen systemet ditt sender UT til klientens URL
type WebhookPayload struct {
	ID      string    `json:"id"`
	Country string    `json:"country"`
	Event   string    `json:"event"`
	Time    time.Time `json:"time"`

	// Disse brukes normalt kun om det er en THRESHOLD event som har truffet
	Feature      string   `json:"feature,omitempty"`
	Operator     string   `json:"operator,omitempty"`
	TriggerValue *float64 `json:"triggerValue,omitempty"`
}
