package api

import (
	"cloud.google.com/go/firestore"
)

type Handler struct {
	DB *firestore.Client
}

// Vi lagrer navnene på database-samlingene (collections) som konstanter.
// Dette gjør vi for å unngå skrivefeil i koden.
const (
	RegistrationsCollection = "registrations"
	NotificationsCollection = "notifications"
)
