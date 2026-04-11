package main

import (
	"envdash/internal/database"
	"log"
	"net/http"
	"os"

	"assignment-2/internal/api"
	"assignment-2/internal/database"
	"assignment-2/internal/models"
)

func main() {
	// 1. Les port fra miljøvariabler (Viktig for Docker/OpenStack deployment)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 2. START DATABASEN:
	// Vi kaller funksjonen vi lagde i firebaseinit.go
	client, err := database.GetFirebaseClient()
	if err != nil {
		log.Fatalf("Klarte ikke starte Firebase: %v", err) // Avslutter programmet hvis DB feiler
	}
	// defer sørger for at forbindelsen til databasen lukkes pent når serveren (main) skrus av.
	defer client.Close()

	// 3. SETT OPP HANDLER (Dependency Injection):
	// Vi gir databaseklienten vår til Handleren. Nå har alle rute-funksjonene tilgang til DB!
	h := &api.Handler{
		DB: client,
	}

	// 4. Sett opp ruting ved hjelp av standardbiblioteket
	mux := http.NewServeMux()

	// Registrations endpoints
	mux.HandleFunc("POST /envdash/v1/registrations/", handlePostRegistration)
	mux.HandleFunc("GET /envdash/v1/registrations/{id}", handleGetRegistration)
	mux.HandleFunc("GET /envdash/v1/registrations/", handleGetAllRegistrations)
	mux.HandleFunc("PUT /envdash/v1/registrations/{id}", handlePutRegistration)
	mux.HandleFunc("DELETE /envdash/v1/registrations/{id}", handleDeleteRegistration)

	// Dashboards endpoint
	mux.HandleFunc("GET /envdash/v1/dashboards/{id}", handleGetDashboard)

	// Notifications endpoints
	mux.HandleFunc("POST /envdash/v1/notifications/", handlePostNotification)
	mux.HandleFunc("GET /envdash/v1/notifications/{id}", handleGetNotification)
	mux.HandleFunc("GET /envdash/v1/notifications/", handleGetAllNotifications)
	mux.HandleFunc("DELETE /envdash/v1/notifications/{id}", handleDeleteNotification)

	// Status endpoint
	mux.HandleFunc("GET /envdash/v1/status/", handleGetStatus)

	log.Printf("Starter server på port %s", port)
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatalf("Kunne ikke starte server: %v", err)
	}
}

// Dummy handlers (Disse bør ligge i internal/api)

func handleGetRegistration(w http.ResponseWriter, r *http.Request)     {}
func handleGetAllRegistrations(w http.ResponseWriter, r *http.Request) {}
func handlePutRegistration(w http.ResponseWriter, r *http.Request)     {}
func handleDeleteRegistration(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
func handleGetDashboard(w http.ResponseWriter, r *http.Request)        {}
func handlePostNotification(w http.ResponseWriter, r *http.Request)    {}
func handleGetNotification(w http.ResponseWriter, r *http.Request)     {}
func handleGetAllNotifications(w http.ResponseWriter, r *http.Request) {}
func handleDeleteNotification(w http.ResponseWriter, r *http.Request)  {}
func handleGetStatus(w http.ResponseWriter, r *http.Request)           {}
