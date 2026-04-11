package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"assignment2/internal/models" // Husk å bytte ut "assignment2" med det vi kalte prosjektet i go.mod
)

// HandlePostRegistration håndterer POST /envdash/v1/registrations/
func (h *Handler) HandlePostRegistration(w http.ResponseWriter, r *http.Request) {
	// frigjøre minne.
	defer r.Body.Close()

	var reg models.Registration

	// json.NewDecoder leser JSON-dataen fra forespørselen (r.Body) og prøver å
	// oversette det ("Decode") til Go-structen vår (reg).
	if err := json.NewDecoder(r.Body).Decode(&reg); err != nil {
		http.Error(w, "Ugyldig JSON-payload", http.StatusBadRequest)
		return
	}

	// Oppgaven krever at vi lagrer når dokumentet sist ble endret.
	reg.LastChange = time.Now().Format("20060102 15:04")

	ctx := r.Context()

	// h.DB.Collection().Add() legger til et nytt dokument i databasen.
	// Firestore vil automatisk generere en unik, tilfeldig ID for oss.
	ref, _, err := h.DB.Collection(RegistrationsCollection).Add(ctx, reg)
	if err != nil {
		log.Printf("Klarte ikke å lagre registrering i DB: %v", err)
		http.Error(w, "Feil ved lagring i database", http.StatusInternalServerError)
		return
	}

	// Vi lagrer den nye databasen-ID-en i structen vår, slik at vi kan sende den tilbake til brukeren.
	reg.ID = ref.ID

	// Forteller klienten som sendte forespørselen at de får JSON tilbake, og at alt gikk bra (201 Created).
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// Gjør Go-map-en vår om til JSON og sender den ut i responskroppen (w).
	json.NewEncoder(w).Encode(map[string]string{
		"id": reg.ID,
	})
}

// HandleGetRegistration håndterer GET /envdash/v1/registrations/{id}
func (h *Handler) HandleGetRegistration(w http.ResponseWriter, r *http.Request) {
	// Trekker ut {id} fra URL-en. Dette fungerer innebygd i Go 1.22 og nyere.
	id := r.PathValue("id")
	ctx := r.Context()

	// Forsøker å hente det ene dokumentet som matcher ID-en.
	doc, err := h.DB.Collection(RegistrationsCollection).Doc(id).Get(ctx)
	if err != nil {
		log.Printf("Feil ved henting av dokument %s: %v", id, err)
		http.Error(w, "Registrering ikke funnet", http.StatusNotFound)
		return
	}

	var reg models.Registration
	// DataTo kopierer dataene fra Firestore-dokumentet over i Go-structen vår.
	if err := doc.DataTo(&reg); err != nil {
		log.Printf("Feil ved parsing av data: %v", err)
		http.Error(w, "Klarte ikke å lese data", http.StatusInternalServerError)
		return
	}

	// Firestore lagrer vanligvis ikke ID-en *inne* i selve dokumentet,
	// så vi må sette den manuelt basert på referansen for at den skal bli med i JSON-responsen.
	reg.ID = doc.Ref.ID

	// Sender structen tilbake til brukeren som JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reg)
}
