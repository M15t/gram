package firebase

import (
	"context"
	"io"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	auth "firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

// Config represents the configuration
type Config struct {
	FirebaseCredentials string // the path to the firebase credentials JSON file
}

// Service represents the firebase service
type Service struct {
	cfg *Config
	ctx context.Context
	app *firebase.App
	auc *auth.Client
	fsc *firestore.Client
}

// New initializes firebase service with default config
func New(cfg *Config) *Service {
	// * Initialize Firebase Admin SDK with the fetched credentials
	// download the content of the JSON key file from the URL
	response, err := http.Get(cfg.FirebaseCredentials)
	if err != nil {
		log.Printf("Error downloading JSON key: %v\n", err)
	}
	defer response.Body.Close()

	// read the content of the response body
	jsonKey, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error reading JSON key content: %v\n", err)
	}

	opt := option.WithCredentialsJSON(jsonKey)
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	// * Initialize Auth client
	auc, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error initializing Auth client: %v\n", err)
	}

	// * Initialize Firestore client
	fsc, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("error initializing Firestore client: %v\n", err)
	}
	// ! gonna monitor if have more issues
	// defer fsc.Close()

	return &Service{
		cfg: cfg,
		ctx: ctx,
		app: app,
		auc: auc,
		fsc: fsc,
	}
}
