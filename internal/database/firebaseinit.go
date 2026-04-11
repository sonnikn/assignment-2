package database

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// ctx (Context) brukes i Go for å håndtere tidsavbrudd (timeouts) og kanselleringer av funksjoner.
// Siden databasekall kan ta tid, krever Firebase at vi sender med en context.
var ctx context.Context

// GetFirebaseContext returnerer en standard context (Background) hvis den ikke allerede eksisterer.
func GetFirebaseContext() context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return ctx
}

// GetFirebaseClient setter opp forbindelsen til Firestore og returnerer klienten.
func GetFirebaseClient() (*firestore.Client, error) {
	ctx = GetFirebaseContext()

	// 1. FINN CREDENTIALS:
	// I skyen (som OpenStack/Heroku) setter vi ofte miljøvariabler i stedet for å laste opp filer.
	// Her sjekker vi først om miljøvariabelen finnes. Hvis ikke, antar vi at vi kjører lokalt
	// og leter etter filen "service-account.json" i rotmappen.
	credFilePath := os.Getenv("FIREBASE_CREDENTIALS")
	if credFilePath == "" {
		credFilePath = "./service-account.json"
	}

	// 2. INITIALISER FIREBASE:
	// Vi forteller Firebase SDK-en hvor passordfilen vår (credentials) ligger.
	sa := option.WithCredentialsFile(credFilePath)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Printf("Feil ved initialisering av Firebase App: %v", err)
		return nil, err
	}

	// 3. OPPRETT FIRESTORE-KLIENT:
	// Vi ber Firebase om å gi oss en spesifikk klient for Firestore-databasen.
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Printf("Feil ved tilkobling til Firestore: %v", err)
		return nil, err
	}

	// Hvis alt gikk bra, returnerer vi klienten.
	return client, nil
}
